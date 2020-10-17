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
    
    func checkAuth(login: String, password: String) -> Promise<Bool>{
        return Promise{ promise in
        var res = false
        network.getJson(url: Constant.authURL).done({ (data) in
            switch data.code{
            case .access:
                promise.fulfill(true)
            case .error:
                promise.fulfill(false)
            }
        }).catch { (error) in
            promise.reject(error)
        }
        
    }
    }
}
