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
        network.getJson(url: Constant.authURL).done({ (json) in
            res = true
            promise.fulfill(true)
        })
        
    }
    }
}
