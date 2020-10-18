//
//  MapViewController.swift
//  AutoSpace
//
//  Created by andarbek on 17.10.2020.
//

import UIKit
import YandexMapsMobile
import PromiseKit
import Alamofire
import SwiftyJSON



class MapViewController: UIViewController, YMKMapCameraListener, YMKMapObjectTapListener, YMKUserLocationObjectListener {
    
    @IBOutlet weak var finishRouteView: UIButton!
    
    @IBOutlet weak var profileViewBut: UIButton!
    @IBOutlet weak var supportViewBut: UIButton!
    @IBOutlet weak var locationViewBut: UIButton!
    @IBOutlet weak var homeViewBut: UIButton!
    var drivingSession: YMKDrivingSession?
    var route: YMKPolylineMapObject?
    
    func onMapObjectTap(with mapObject: YMKMapObject, point: YMKPoint) -> Bool {
        let alert = UIAlertController(title: "Адрес", message: self.info?.user?.address, preferredStyle: .alert);
        alert.view.backgroundColor = UIColor.darkGray;
        alert.view.alpha = 0.8;
        alert.view.layer.cornerRadius = 15;

        self.present(alert, animated: true);
        DispatchQueue.main.asyncAfter(deadline: DispatchTime.now() + 2) {
            alert.dismiss(animated: true);
        }
        return true
    }
    
    @IBAction func finishRoute(_ sender: UIButton) {
        self.mapView?.mapWindow.map.mapObjects.remove(with: self.route!)
        self.finishRouteView.isHidden = true
    }
    
    @IBOutlet weak var mapView: YMKMapView?
    
    var info : ASRequest?
    
    let mapKit = YMKMapKit.sharedInstance()

    lazy var userLocationLayer = YMKMapKit.sharedInstance().createUserLocationLayer(with: (mapView?.mapWindow)!)
    var position = YMKCameraPosition()
    var isLaunched = false
    
    private var circleMapObjectTapListener: YMKMapObjectTapListener!
    private var animationIsActive = true
    override func viewDidLoad() {
        super.viewDidLoad()
        finishRouteView.isHidden = true
        finishRouteView.titleLabel?.textAlignment = .center
        finishRouteView.layer.cornerRadius = 5
        finishRouteView.layer.borderWidth = 1
        finishRouteView.layer.borderColor = UIColor.black.cgColor
        circleMapObjectTapListener = CircleMapObjectTapListener(controller: self);

        mapView?.mapWindow.map.move(with:
            YMKCameraPosition(target: YMKPoint(latitude: 0, longitude: 0), zoom: 14, azimuth: 0, tilt: 0))
        
        let scale = UIScreen.main.scale

        userLocationLayer.setVisibleWithOn(true)
        userLocationLayer.isHeadingEnabled = true
        userLocationLayer.setObjectListenerWith(self)
        
        
        userLocationLayer.setAnchorWithAnchorNormal(CGPoint(x: 0.5 * (mapView?.frame.size.width)! * scale, y: 0.5 * (mapView?.frame.size.height)! * scale), anchorCourse: CGPoint(x: 0.5 * (mapView?.frame.size.width)! * scale, y: 0.83 * (mapView?.frame.size.height)! * scale))
        userLocationLayer.setObjectListenerWith(self)
        mapView?.mapWindow.map.addCameraListener(with: self)
        guard let latitude = self.info?.user?.addressCoord.latitude else {return}
        guard let longitude = self.info?.user?.addressCoord.longitude else {return}
        let point = YMKPoint(latitude: Double(latitude), longitude: Double(longitude))
        let image = UIImage(named: "marker")
        let place = mapView?.mapWindow.map.mapObjects.addPlacemark(with: point, image: image!)
        place?.addTapListener(with: self)
        getSpots()
    }
    
    override func viewWillDisappear(_ animated: Bool) {
        super.viewWillDisappear(animated)
        self.animationIsActive = false
    }
 
    
    

    @IBAction func openSupport(_ sender: UIButton) {
        let storyBoard : UIStoryboard = UIStoryboard(name: "Map", bundle:nil)
        let nextViewController = storyBoard.instantiateViewController(withIdentifier: "SupportViewController") as! SupportViewController
        nextViewController.modalPresentationStyle = .popover
        nextViewController.email = info?.user?.email
        self.present(nextViewController, animated:true, completion:nil)
    }
    
    @IBAction func openProfile(_ sender: UIButton) {
        performSegue(withIdentifier: "showProfile", sender: UIButton.self)
    }
    
    override func prepare(for segue: UIStoryboardSegue, sender: Any?) {
        if segue.identifier == "showProfile" {
            if let vc = segue.destination as? ProfileViewController {
                vc.info = self.info
            }
        }
    }
    
    @IBAction func findMe(_ sender: UIButton) {
        self.position = self.userLocationLayer.cameraPosition() ?? YMKCameraPosition()
        mapView?.mapWindow.map.move(with: self.position)
    }
    
    @IBAction func findHome(_ sender: UIButton) {
        guard let latitude = self.info?.user?.addressCoord.latitude else {return}
        guard let longitude = self.info?.user?.addressCoord.longitude else {return}
        let point = YMKPoint(latitude: Double(latitude), longitude: Double(longitude))
        let image = UIImage(named: "marker")
        let place = mapView?.mapWindow.map.mapObjects.addPlacemark(with: point, image: image!)
        place?.addTapListener(with: self)
        mapView?.mapWindow.map.move(with: YMKCameraPosition(target: point, zoom: 18, azimuth: 120, tilt: 100))
        
        self.position = self.userLocationLayer.cameraPosition() ?? YMKCameraPosition()
        self.userLocationLayer.resetAnchor()

    }
    
    
    func addPlacemarkListener(withPlacemarkListener placemarkListener: YMKSearchLayerTapHandler) {

    }
    
    func onCameraPositionChanged(with map: YMKMap,
                                     cameraPosition: YMKCameraPosition,
                                     cameraUpdateReason: YMKCameraUpdateReason,
                                     finished: Bool) {
        if finished == false {
            self.position = self.userLocationLayer.cameraPosition() ?? YMKCameraPosition()
            self.userLocationLayer.resetAnchor()
        }
        }
    
    
    func onObjectAdded(with view: YMKUserLocationView) {
        guard let image = UIImage(named:"UserArrow") else {return}
        view.arrow.setIconWith(image)
        
        let pinPlacemark = view.pin.useCompositeIcon()
        
        pinPlacemark.setIconWithName("icon",
            image: UIImage(named:"Icon")!,
            style:YMKIconStyle(
                anchor: CGPoint(x: 0, y: 0) as NSValue,
                rotationType:YMKRotationType.rotate.rawValue as NSNumber,
                zIndex: 0,
                flat: true,
                visible: true,
                scale: 1.5,
                tappableArea: nil))
        
        pinPlacemark.setIconWithName(
            "pin",
            image: UIImage(named:"SearchResult")!,
            style:YMKIconStyle(
                anchor: CGPoint(x: 0.5, y: 0.5) as NSValue,
                rotationType:YMKRotationType.rotate.rawValue as NSNumber,
                zIndex: 1,
                flat: true,
                visible: true,
                scale: 1,
                tappableArea: nil))

        view.accuracyCircle.fillColor = UIColor.blue
    }
    
    func onObjectRemoved(with view: YMKUserLocationView) {
    }
    
    func onObjectUpdated(with view: YMKUserLocationView, event: YMKObjectEvent) {
//        self.getSpots()
    }
    
    
    
    private class CircleMapObjectTapListener: NSObject, YMKMapObjectTapListener {
        private weak var controller: MapViewController?

        init(controller: MapViewController) {
            self.controller = controller
        }

        func onMapObjectTap(with mapObject: YMKMapObject, point: YMKPoint) -> Bool {
            print("2134564321")
            if let circle = mapObject as? YMKCircleMapObject {
                if let userData = circle.userData as? CircleMapObjectUserData {
                    var mess = ""
                    if userData.free == 0 {
                        mess = "Все парковочные места заняты"
                    }
                    else {
                        mess = "Свободно \(userData.free) из \(userData.total) мест"
                    }
                    let alert = UIAlertController(title: "Информация о парковке:", message: mess, preferredStyle: .alert)
                    alert.addAction(UIAlertAction(title: "Закрыть", style: .cancel, handler: nil))
                    if userData.free != 0 {
                        alert.addAction(UIAlertAction(title: "Маршрут", style: .default, handler: { [weak self]action in self!.controller?.getRoute(point: YMKPoint(latitude: userData.coor.latitude, longitude: userData.coor.longitude))}))
                    }
                    
                    controller?.present(alert, animated: true);

                }
            }
            return true;
        }
    }

    private class CircleMapObjectUserData {
        let id: Int32;
        let free: Int32;
        let total: Int32;
        let coor: Coordinate
        init(id: Int32, free: Int32, total: Int32, coor: Coordinate) {
            self.id = id;
            self.total = total;
            self.free = free;
            self.coor = coor;
        }
    }

    func createTappableCircle(free: Int32, total: Int32, point: YMKPoint) -> YMKCircleMapObject {
        let mapObjects = mapView?.mapWindow.map.mapObjects;
        var color = UIColor.green.withAlphaComponent(0.3)
        let coor = Coordinate(latitude: point.latitude, longitude: point.longitude)
        if free == 0 {
            color = UIColor.red.withAlphaComponent(0.3)
 
        }
        let circle = mapObjects!.addCircle (
            with: YMKCircle(center: point, radius: 15),
            stroke: color,
            strokeWidth: 1,
            fill: color)
        circle.zIndex = 100
        let image = UIImage(named: "spot")
        let parking = self.mapView?.mapWindow.map.mapObjects.addPlacemark(with: point, image: image!)
        parking?.zIndex = 100
        circle.userData = CircleMapObjectUserData(id: 1, free: free, total: total, coor: coor);
        return circle
    }
    
    
    func getRoute(point: YMKPoint) {
        let startPoint = YMKPoint(latitude: self.position.target.latitude, longitude: self.position.target.longitude)
        let requestPoints : [YMKRequestPoint] = [
            YMKRequestPoint(point: startPoint, type: .waypoint, pointContext: nil),
                    YMKRequestPoint(point: point, type: .waypoint, pointContext: nil),
                    ]
                
        let responseHandler = {(routesResponse: [YMKDrivingRoute]?, error: Error?) -> Void in
            if let routes = routesResponse {
                self.onRoutesReceived(routes)
            } else {
                self.onRoutesError(error!)
            }
        }
        
        let drivingRouter = YMKDirections.sharedInstance().createDrivingRouter()
        drivingSession = drivingRouter.requestRoutes(
            with: requestPoints,
            drivingOptions: YMKDrivingDrivingOptions(),
            vehicleOptions: YMKDrivingVehicleOptions(),
            routeHandler: responseHandler)
        self.finishRouteView.isHidden = false
    }
            
    func onRoutesReceived(_ routes: [YMKDrivingRoute]) {
        let mapObjects = mapView?.mapWindow.map.mapObjects
        if routes.count > 0 {
            self.route = mapObjects!.addPolyline(with: routes[0].geometry)
        }
    }
            
    func onRoutesError(_ error: Error) {
        let routingError = (error as NSError).userInfo[YRTUnderlyingErrorKey] as! YRTError
        var errorMessage = "Unknown error"
        if routingError.isKind(of: YRTNetworkError.self) {
            errorMessage = "Network error"
        } else if routingError.isKind(of: YRTRemoteError.self) {
            errorMessage = "Remote server error"
        }
        
        let alert = UIAlertController(title: "Error", message: errorMessage, preferredStyle: .alert)
        alert.addAction(UIAlertAction(title: "OK", style: .default, handler: nil))
        
        present(alert, animated: true, completion: nil)
    }
    
    func getSpots(){
        AF.request(Constant.getSpotsURL, method: .post).responseJSON(completionHandler: { [weak self] response in
            do{
                if (0200...299).contains(response.response?.statusCode ?? 0){
                    let json = try JSON(data: response.data!)
                    for i in json["spaces"]{
                        
                        self!.createTappableCircle(free: i.1["free"].int32!, total: i.1["total"].int32!, point: YMKPoint(latitude: i.1["lat"].double!, longitude: i.1["long"].double!)).addTapListener(with: self!.circleMapObjectTapListener);
                    }
                    
                }

            } catch {
                print("Error")
            }
            
        }).resume()
    
    }
    
}
