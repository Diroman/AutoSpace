//
//  NetworkManager.swift
//  AutoSpace
//
//  Created by andarbek on 17.10.2020.
//

import Foundation
import Alamofire
import PromiseKit

class NetworkManager {
    
    func performRequest(_ request: NetworkRequest, validStatusCodes: [Int] = (200...299).map({$0})) -> Promise<Data?> {
        return Promise { [weak self] promise in
            print("2 \(promise)")
            AF.request(request.url,
                       method: request.method)
                .responseData(completionHandler: { (response) in
                    if validStatusCodes.contains(response.response?.statusCode ?? 0) {
                        promise.fulfill(response.data)
                    } else {
                        print(response.data)
                        print(response.value)
                        print(response.response)
                    }
                    
                })
            //requestHandler?(request)
        }
    }
    
}
