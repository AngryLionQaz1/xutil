package com.snow.golang.common.repository;

import com.snow.golang.common.pojo.User;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.Optional;


public interface UserRepository extends JpaRepository<User,Long> {



    Optional<User> findByUsername(String name);
}
