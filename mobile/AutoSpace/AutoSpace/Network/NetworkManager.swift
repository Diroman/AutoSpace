//
//  NetworkManager.swift
//  AutoSpace
//
//  Created by andarbek on 17.10.2020.
//

import Foundation
import Alamofire
import PromiseKit
import SwiftyJSON

class NetworkManager {
    
    func performRequest(_ request: NetworkRequest, validStatusCodes: [Int] = (200...299).map({$0})) -> Promise<ASRequest> {
        return Promise { [weak self] promise in
            print("2 \(promise)")
            AF.request(request.url,
                       method: request.method, parameters: Constant.parameters, encoding: JSONEncoding.default)
                .responseData(completionHandler: { (response) in
                    if validStatusCodes.contains(response.response?.statusCode ?? 0) {
                        print(response.response)
                        let res = ASRequest(code: .access, data: response.data)
                        promise.fulfill(res)
                    } else {
                        let error = JSON(response.response?.statusCode)
                        let errType: ErrorType?
                        switch error {
                        case 400:
                            errType = .wrongPassword
                        case 401:
                            errType = .notRegistered
                        default:
                            errType = .unknown
                        }

                        print(error)
                        promise.fulfill(ASRequest(code: .error, error: errType))
                    }
                    
                })
            //requestHandler?(request)
        }
    }
    
}
