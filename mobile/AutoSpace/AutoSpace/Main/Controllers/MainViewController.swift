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
        spinner.hidesWhenStopped = true
        loadingBackground.isHidden = true
        self.hideKeyboardWhenTappedAround()
        
    }
    
    override func viewWillLayoutSubviews() {
        super.viewWillLayoutSubviews()
//        login.textColor = .white
//        login.tintColor = .white
        login.borderStyle = .roundedRect
        password.borderStyle = .roundedRect
//        password.textColor = .white
//        password.tintColor = .white
        enterButtonView.layer.cornerRadius = 10
        enterButtonView.layer.borderWidth = 1
    }
    
    @IBAction func enter(_ sender: UIButton) {
        spinner.startAnimating()
        enterButtonView.isEnabled = false
        loadingBackground.isHidden = false
        
        self.model.checkAuth(login: login.text!, password: password.text!).done { [weak self] (flag) in
            self?.enterButtonView.isEnabled = true
            self?.loadingBackground.isHidden = true
            self?.spinner.stopAnimating()
            var errMess = ""
            switch flag.code{
            case .access:
                
                let storyBoard : UIStoryboard = UIStoryboard(name: "Map", bundle:nil)
                let nextViewController = storyBoard.instantiateViewController(withIdentifier: "MapViewController") as! MapViewController
                nextViewController.modalPresentationStyle = .fullScreen
                nextViewController.info = flag
                self?.present(nextViewController, animated:true, completion:nil)
            case .error:
                
                switch flag.error {
                case .wrongPassword:
                    errMess = "Неправильный логин или пароль"
                case .notRegistered:
                    errMess = "Пользователь не зарегистрирован"
                case .unknown:
                    errMess = "Произошла ошибка, повторите снова"
                default:
                    errMess = "Произошла ошибка, повторите снова"
                }
                let alert = UIAlertController(title: "Ошибка", message: errMess, preferredStyle: .alert)
                let ok = UIAlertAction(title: "OK", style: .default, handler: nil)
                alert.addAction(ok)
                self!.present(alert, animated: true, completion: nil)
                
            case .none:
                let alert = UIAlertController(title: "Ошибка", message: errMess, preferredStyle: .alert)
                let ok = UIAlertAction(title: "OK", style: .default, handler: nil)
                alert.addAction(ok)
                self!.present(alert, animated: true, completion: nil)
            }
            //сделать кнопку активной
        }
            
        }
    
    }



