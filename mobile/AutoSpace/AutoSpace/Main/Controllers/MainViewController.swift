//
//  ViewController.swift
//  AutoSpace
//
//  Created by andarbek on 16.10.2020.
//

import UIKit
import PromiseKit

class MainViewController: UIViewController {

    @IBOutlet weak var backgroundView: UIView!
    @IBOutlet weak var login: UITextField!
    @IBOutlet weak var password: UITextField!
    @IBOutlet weak var loadingBackground: UIView!
    @IBOutlet weak var spinner: UIActivityIndicatorView!
    @IBOutlet weak var enterButtonView: UIButton!
    
    let model = MainModel()
    
    override func viewDidLoad() {
        super.viewDidLoad()
        backgroundView.tintColor = #colorLiteral(red: 0.8039215803, green: 0.8039215803, blue: 0.8039215803, alpha: 1)
        spinner.hidesWhenStopped = true
        loadingBackground.isHidden = true
    }
    
    @IBAction func enter(_ sender: UIButton) {
        spinner.startAnimating()
        enterButtonView.isSelected = true
        self.model.checkAuth(login: login.text!, password: password.text!).done { (flag) in
            switch flag{
            case true:
                let view = MapViewController(nibName: "MapViewController", bundle: nil)
                self.present(view, animated: true, completion: nil)
                
            case false:
                print("false")
                //обработать случаи, когда присылается дичь
                //алерт
            }
            //сделать кнопку активной
        }
            
        }
    }



