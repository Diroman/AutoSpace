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
            //переход
            //сделать кнопку активной
            
            //обработать случаи, когда присылается дичь
        }
            
        }
    }



