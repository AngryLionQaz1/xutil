/**
 * Copyright (C), 2015-2019, 重庆了赢科技有限公司
 * FileName: UserRedis
 * Author:   萧毅
 * Date:     2019/2/26 14:07
 * Description:
 */
package com.snow.golang.common.repository;


import com.snow.golang.common.pojo.User;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cache.annotation.Cacheable;
import org.springframework.stereotype.Component;

import java.util.Optional;

@Component
public class UserRedis {

    @Autowired
    private UserRepository userRepository;

    @Cacheable(key="'user_'+#id",value="'user")
    public User getUser(Long userId){
        Optional<User> byUsername = userRepository.findById(userId);
        if (!byUsername.isPresent())return null;
        User user=byUsername.get();
        return user;
    }





}