package com.snow.golang.common.pojo;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.*;
import java.util.List;

@Builder
@Entity
@Data
@Table(name = "s_authority")
@AllArgsConstructor
@NoArgsConstructor
public class Authority {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    /**权限名*/
    private String name;
    /**权限编码*/
    @Column(unique = true)
    private String code;
    /**父级编码*/
    private String pCode;
    /**类别名*/
    private String pName;
    /**权限uri*/
    private String uri;
    /**详细描述*/
    private String details;
    /**顶级 true 为顶级 false 非顶级*/
    private Boolean flag;
    /**是不是菜单*/
    private Boolean ifMenu;
    /**排序*/
    private Integer mOrder;
    @Transient
    private List<Authority> authorities;











}

