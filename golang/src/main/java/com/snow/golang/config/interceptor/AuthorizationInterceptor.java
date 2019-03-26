package com.snow.golang.config.interceptor;

import com.alibaba.fastjson.JSON;
import com.snow.golang.common.bean.Config;
import com.snow.golang.common.pojo.User;
import com.snow.golang.common.repository.UserRepository;
import com.snow.golang.config.security.SecurityContextHolder;
import com.snow.golang.config.token.JWTToken;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.MediaType;
import org.springframework.stereotype.Component;
import org.springframework.web.method.HandlerMethod;
import org.springframework.web.servlet.ModelAndView;
import org.springframework.web.servlet.handler.HandlerInterceptorAdapter;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.io.PrintWriter;
import java.util.Optional;

import static com.snow.golang.common.bean.Result.over;



/**
 * 自定义拦截器，判断此次请求是否有权限
 *
 *
 */
@Component
public class AuthorizationInterceptor extends HandlerInterceptorAdapter {


    @Autowired
    private Config config;
    @Autowired
    private JWTToken jwtToken;
    @Autowired
    private SecurityContextHolder securityContextHolder;
    @Autowired
    private UserRepository userRepository;



    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) {
        //如果不是映射到方法直接通过
        if (!(handler instanceof HandlerMethod)) return true;
        String token=request.getHeader(config.getAuthorization());
        if (Optional.ofNullable(token).isPresent()&& jwtToken.validateToken(token)&&jwtToken.getUserId(token)!=null){
            Optional<User> o=userRepository.findById(Long.valueOf(jwtToken.getUserId(token)));
            if (o.isPresent())return true;
        }
        //如果验证token失败，返回错误信息
        response(response);
        return false;
    }

    @Override
    //在后端控制器执行后调用
    public void postHandle(HttpServletRequest request, HttpServletResponse response, Object handler, ModelAndView modelAndView) throws Exception {
        super.postHandle(request, response, handler, modelAndView);
    }

    @Override
    //整个请求执行完成后调用
    public void afterCompletion(HttpServletRequest request, HttpServletResponse response, Object handler, Exception ex) throws Exception {
        super.afterCompletion(request, response, handler, ex);
        securityContextHolder.removeUser();
    }




    
    //返回错误信息
    public void response(HttpServletResponse response){
        response.setHeader("Cache-Control", "no-store");
        response.setHeader("Pragma", "no-cache");
        response.setContentType(MediaType.APPLICATION_JSON_VALUE);
        response.setCharacterEncoding("UTF-8");
        PrintWriter out= null;
        try {
            out = response.getWriter();
            out.write(JSON.toJSONString(over()));
            out.flush();
        } catch (IOException e) {
            e.printStackTrace();
        }finally {
            out.close();
        }

    }


}


