//
//  Constant.swift
//  AutoSpace
//
//  Created by andarbek on 16.10.2020.
//

import Foundation
import UIKit

class Constant{
    static let backgroundColor = #colorLiteral(red: 0.8039215803, green: 0.8039215803, blue: 0.8039215803, alpha: 1)
    
    //static let authURL = "http://jsonplaceholder.typicode.com/posts"
    
    static let authURL = "http://192.168.31.44:8080/login"
    static let sendCommentURL = "http://192.168.31.44:8080/send-email"
    static let getSpotsURL = "http://192.168.31.44:8080/get-free-space"
    
    static var parameters: [String: String] = [
        "login" : "" ,
        "password" : ""
    ]
    
    static var sendParam: [String: Any] = [
        "error_code" : 92 ,
        "comment" : "huhu",
        "email" : "hhaa"
    ]
    
}
