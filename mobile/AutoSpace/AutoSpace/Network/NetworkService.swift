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
    
    func getJson(url: String) -> Promise<JSON> {
        return Promise { [unowned self] promise in
            let request = NetworkRequest.create(url: url, method: .get)
            print("request")
            print(request.url)
            self.performRequest(request).done({ (data) in
                print("1 \(data)")
                guard let data = data else { throw NSError() }
                guard let json: JSON = try? JSON(data: data) else { throw NSError() }
                print(json)
                promise.fulfill(json)
            }).catch({ (error) in
                promise.reject(error)
            })
        }
    }
}
