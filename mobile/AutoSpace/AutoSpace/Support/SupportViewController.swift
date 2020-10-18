//
//  SupportViewController.swift
//  AutoSpace
//
//  Created by andarbek on 18.10.2020.
//

import UIKit
import PromiseKit

class SupportViewController: UIViewController, UIPickerViewDelegate, UIPickerViewDataSource  {
    
    let thePicker = UIPickerView()
    let pickerData = ["Ошибка в приложении", "Cвободное место занято", "Неправильная парковка"]
    let pickerDict = ["Ошибка в приложении" : 800, "Cвободное место занято" : 801, "Неправильная парковка" : 802]
    
    
    @IBOutlet weak var errorTf: UITextField!
    @IBOutlet weak var commentsTf: UITextView!
    @IBOutlet weak var sendButView: UIButton!
    @IBOutlet weak var spinner: UIActivityIndicatorView!
    
    let model = SupportModel()
    var email : String?
    
    override func viewDidLoad() {
        super.viewDidLoad()
        commentsTf.backgroundColor = .white
        spinner.hidesWhenStopped = true
        self.hideKeyboardWhenTappedAround()
        thePicker.delegate = self
        errorTf.inputView = thePicker
    }
    
    override func viewWillLayoutSubviews() {
        super.viewWillLayoutSubviews()
        commentsTf.layer.cornerRadius = 10
        commentsTf.layer.borderWidth = 1
        errorTf.borderStyle = .roundedRect
        sendButView.layer.cornerRadius = 10
        sendButView.layer.borderWidth = 1
    }
    

    @IBAction func sendComment(_ sender: Any) {
        if errorTf.text != "" && commentsTf.text != "" && commentsTf.text.count < 1001 && email != nil{
            spinner.startAnimating()
            self.sendButView.isEnabled = false
            var message = ""
            var title = ""
            self.model.sendMess(errorCode: pickerDict[errorTf.text!] ?? 800, comment: commentsTf.text!, email: email ?? "").done{ [weak self] flag in
                self?.spinner.stopAnimating()
                self?.sendButView.isEnabled = true
                switch flag.code{
                case .access:
                    message = "Сообщение отправлено!\nСпасибо за обращение!"
                    title = "Готово!"
                case .error:
                    message = "Возникли неполадки! Попробуйте снова."
                    title = "Ошибка"
                case .none:
                    message = "Возникли неполадки! Попробуйте снова."
                    title = "Ошибка"
                }
                let alert = UIAlertController(title: title, message: message, preferredStyle: .alert)
                let ok = UIAlertAction(title: "OK", style: .default, handler: nil)
                alert.addAction(ok)
                self?.present(alert, animated: true, completion: nil)
            }
        }
    }
    
    func numberOfComponents(in pickerView: UIPickerView) -> Int {
        return 1
    }

    func pickerView(_ pickerView: UIPickerView, numberOfRowsInComponent component: Int) -> Int {
        return pickerData.count
    }

    func pickerView( _ pickerView: UIPickerView, titleForRow row: Int, forComponent component: Int) -> String? {
        return pickerData[row]
    }

    func pickerView( _ pickerView: UIPickerView, didSelectRow row: Int, inComponent component: Int) {
        errorTf.text = pickerData[row]
    }
    

}
