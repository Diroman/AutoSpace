//
//  SupportModel.swift
//  AutoSpace
//
//  Created by andarbek on 18.10.2020.
//
import Foundation
import PromiseKit

class SupportModel {
    
    let network = NetworkService()
    
    func sendMess(errorCode: Int, comment: String, email: String) -> Promise<ASRequest>{
        return Promise{ promise in
        //var res = false
            Constant.sendParam["error_code"] = errorCode
            Constant.sendParam["comment"] = comment
            Constant.sendParam["email"] = email
            print(Constant.sendParam["comment"])
            network.sendComment(url: Constant.sendCommentURL).done({ (data) in
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
