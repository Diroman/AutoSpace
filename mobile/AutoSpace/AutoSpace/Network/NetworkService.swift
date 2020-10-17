//
//  NetworkService.swift
//  AutoSpace
//
//  Created by andarbek on 17.10.2020.
//

import Foundation
import Alamofire
import PromiseKit
import SwiftyJSON

class NetworkService: NetworkManager {
    
    func getJson(url: String) -> Promise<ASRequest> {
        return Promise { [unowned self] promise in
            let request = NetworkRequest.login(url: url, method: .post, parameters: Constant.parameters)
            self.performRequest(request).done({ (data) in
                switch data.code{
                case .access:
                    print(data.data)
                    parseJson()
                    promise.fulfill(data)
                case .error:
                    promise.fulfill(data)
                }
               // promise.fulfill(json)
            })
            
        }
    }
    
    func parseJson(){
        
    }
}
