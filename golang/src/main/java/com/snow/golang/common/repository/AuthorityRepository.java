package com.snow.golang.common.repository;


import com.snow.golang.common.pojo.Authority;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.Optional;

public interface AuthorityRepository extends JpaRepository<Authority,Long> {


 Optional<Authority> findByUri(String url);
}

