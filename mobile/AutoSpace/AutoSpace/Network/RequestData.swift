//
//  File.swift
//  AutoSpace
//
//  Created by andarbek on 17.10.2020.
//

import Foundation

enum RequestError: Error {
    case access
    case wrongPassword
    case serverError
}

struct RequestData: Error{
    var error: Data?
    
    init(error: Data) {
        self.error = error
    }
}
