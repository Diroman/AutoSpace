//
//  User.swift
//  AutoSpace
//
//  Created by andarbek on 16.10.2020.
//

import Foundation

struct Address{
    
}


class User {
    var name: String
    var homeAddress: Address
    
    init(name: String, homeAddress: Address){
        self.name = name
        self.homeAddress = homeAddress
    }
}
