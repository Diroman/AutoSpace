//
//  Request.swift
//  AutoSpace
//
//  Created by andarbek on 17.10.2020.
//

import Foundation
import Alamofire

class NetworkRequest{
    
    var url: String!
    var headers: [String: String]!
    var parameters: [String: Any]!
    var method: Alamofire.HTTPMethod!
    
    required init(url: String, method: Alamofire.HTTPMethod) {
        self.url = url
        self.headers = [String: String]()
        self.parameters = [String: Any]()
        self.method = method
    }
    
    
    
    class func create(url: String, method: Alamofire.HTTPMethod) -> Self {
        let request = self.init(url: url, method: method)
        //request.withHeader(headers)
        return request
    }
    
    class func createWithHeader(url: String, method: Alamofire.HTTPMethod, headers: [String: String]) -> Self {
        let request = self.init(url: url, method: method)
        request.withHeader(headers: headers)
        return request
    }
    
    @discardableResult
    func withHeader(headers: [String: String]) -> Self {
        self.headers = headers
        return self
    }
    
    
}
