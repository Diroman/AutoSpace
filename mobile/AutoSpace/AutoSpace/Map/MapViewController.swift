//
//  MapViewController.swift
//  AutoSpace
//
//  Created by andarbek on 17.10.2020.
//

import UIKit
import YandexMapsMobile


class MapViewController: UIViewController, YMKMapCameraListener, YMKUserLocationObjectListener {
    
    @IBOutlet weak var mapView: YMKMapView?
    
    //let TARGET_LOCATION = YMKPoint(latitude: 59.945933, longitude: 30.320045)
    let mapKit = YMKMapKit.sharedInstance()

    lazy var userLocationLayer = YMKMapKit.sharedInstance().createUserLocationLayer(with: (mapView?.mapWindow)!)
    var position = YMKCameraPosition()
    var isLaunched = false
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
       // mapView?.mapWindow.map.move(
//            with: YMKCameraPosition(target: TARGET_LOCATION, zoom: 15, azimuth: 0, tilt: 0),
//            animationType: YMKAnimation(type: YMKAnimationType.smooth, duration: 5),
//            cameraCallback: nil)
        
        mapView?.mapWindow.map.move(with:
            YMKCameraPosition(target: YMKPoint(latitude: 0, longitude: 0), zoom: 14, azimuth: 0, tilt: 0))
        
        let scale = UIScreen.main.scale

        userLocationLayer.setVisibleWithOn(true)
        userLocationLayer.isHeadingEnabled = true
        userLocationLayer.setObjectListenerWith(self)
        
        
        userLocationLayer.setAnchorWithAnchorNormal(CGPoint(x: 0.5 * (mapView?.frame.size.width)! * scale, y: 0.5 * (mapView?.frame.size.height)! * scale), anchorCourse: CGPoint(x: 0.5 * (mapView?.frame.size.width)! * scale, y: 0.83 * (mapView?.frame.size.height)! * scale))
        userLocationLayer.setObjectListenerWith(self)
        mapView?.mapWindow.map.addCameraListener(with: self)
        
//        mapView?.mapWindow.map.addTapListener(with: self)
    }

    
    @IBAction func findMe(_ sender: UIButton) {
        self.position = self.userLocationLayer.cameraPosition() ?? YMKCameraPosition()
        mapView?.mapWindow.map.move(with: self.position)
    }
    
    @IBAction func findHome(_ sender: UIButton) {
        let mapKit = YMKMapKit.sharedInstance()
        
        
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
        print("jhg")
    }
    
    func onObjectUpdated(with view: YMKUserLocationView, event: YMKObjectEvent) {
        print("jh")
    }
    
}
