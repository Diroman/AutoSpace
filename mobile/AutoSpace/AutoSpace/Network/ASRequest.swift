//
//  File.swift
//  AutoSpace
//
//  Created by andarbek on 17.10.2020.
//

import Foundation
import SwiftyJSON

enum ASRequestType {
    case access
    case error
}

enum ErrorType{
    case wrongPassword
    case notRegistered
    case unknown
}

struct ASRequest {
    var code : ASRequestType
    var data : Data?
    var error : ErrorType?
    var json : JSON?
}
