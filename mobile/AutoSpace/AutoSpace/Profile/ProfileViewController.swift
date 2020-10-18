//
//  ProfileViewController.swift
//  AutoSpace
//
//  Created by andarbek on 17.10.2020.
//

import UIKit

class ProfileViewController: UIViewController {

    var info : ASRequest?
    
    @IBOutlet weak var pointsView: UIView!
    @IBOutlet weak var fio: UILabel!
    @IBOutlet weak var address: UILabel!
    @IBOutlet weak var phone: UILabel!
    
    @IBOutlet weak var email: UILabel!
    @IBOutlet weak var points: UILabel!
    @IBOutlet weak var circleView: UIView!
    
    override func viewDidLoad() {
        super.viewDidLoad()

        setInfo()
    }
    
    func setInfo(){
        self.fio.text = self.info?.user?.username
        self.address.text = self.info?.user?.address
        self.phone.text = self.info?.user?.mobile
        self.email.text = self.info?.user?.email
        self.points.text = "69"
    }
    
    override func viewWillLayoutSubviews() {
        super.viewWillLayoutSubviews()
        self.circleView.layer.cornerRadius = 45
        self.circleView.layer.masksToBounds = true;
        self.pointsView.layer.cornerRadius = 10
        self.pointsView.layer.borderWidth = 1
    }
    


}
