//
//  MainModel.swift
//  AutoSpace
//
//  Created by andarbek on 17.10.2020.
//

import Foundation
import PromiseKit

class MainModel {
    
    let network = NetworkService()
    
    func checkAuth(login: String, password: String) -> Promise<ASRequest>{
        return Promise{ promise in
        //var res = false
        Constant.parameters["login"] = login
        Constant.parameters["password"] = password
        network.getJson(url: Constant.authURL).done({ (data) in
            switch data.code{
            case .access:
                promise.fulfill(data)
            case .error:
                promise.fulfill(data)
            case .none:
                promise.fulfill(data)
            }
        }).catch { (error) in
            promise.reject(error)
        }
        
    }
    }
}
