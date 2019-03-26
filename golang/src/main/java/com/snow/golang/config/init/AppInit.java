package com.snow.golang.config.init;


import com.snow.golang.common.bean.Config;
import com.snow.golang.common.pojo.Authority;
import com.snow.golang.common.repository.AuthorityRepository;
import com.snow.golang.config.annotation.Security;
import io.swagger.annotations.ApiOperation;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.ApplicationArguments;
import org.springframework.boot.ApplicationRunner;
import org.springframework.stereotype.Component;
import org.springframework.web.context.WebApplicationContext;
import org.springframework.web.method.HandlerMethod;
import org.springframework.web.servlet.mvc.method.RequestMappingInfo;
import org.springframework.web.servlet.mvc.method.annotation.RequestMappingHandlerMapping;

import java.lang.reflect.Method;
import java.util.*;
import java.util.stream.Collectors;

@Component
public class AppInit implements ApplicationRunner {

    @Autowired
    private WebApplicationContext applicationContext;
    @Autowired
    private AuthorityRepository authorityRepository;
    @Autowired
    private Config config;

    @Override
    public void run(ApplicationArguments args) {
        if (config.getAuthorityInit())new Thread(()->init()).run();
    }


    private void init(){
        RequestMappingHandlerMapping mapping = applicationContext.getBean(RequestMappingHandlerMapping.class);
        //获取url与类和方法的对应信息
        Map<RequestMappingInfo, HandlerMethod> map = mapping.getHandlerMethods();
        List<Authority> authorities = new ArrayList<>();
        int count=0;
        for (RequestMappingInfo info : map.keySet()){
            Method m = map.get(info).getMethod();
            List<Authority> list=menu(count,info,m);
            if (list==null)continue;
            list.stream()
                    .filter(Objects::nonNull)
                    .forEach(i->apiOperation(i,info,m));
            authorities.addAll(list);
            count++;
        }
        authorities = authorities.stream()
                .filter(Objects::nonNull)
                .collect(Collectors.collectingAndThen(
                        Collectors.toCollection(() -> new TreeSet<>(Comparator.comparing(Authority::getCode))), ArrayList::new))
                .stream()
                .filter(i->!checkData(i.getUri()))
                .collect(Collectors.toList());
        if (authorities.size()>0)
        authorityRepository.saveAll(authorities);
    }

    /**
     * 查看数据库中是否存在
     * @param url
     * @return
     */
    private boolean checkData(String url){
        if ("".equals(url)||url==null)return false;
        return authorityRepository.findByUri(url).isPresent();
    }


    /**
     * 获取菜单
     */
    public List<Authority> menu(Integer count, RequestMappingInfo info, Method m){
        //获取类上security
        Security clzss=m.getDeclaringClass().getAnnotation(Security.class);
        if (clzss==null)return null;
        List<Authority> list=new ArrayList<>();
        list.add(menuTop(clzss));
        //获取方法上的security
        Security security = m.getAnnotation(Security.class);
        if (security!=null)list.addAll(menus(clzss,security,count));
        if (security==null)list.add(security(clzss,info,count));
        return list;

    }


    /**
     * 权限
     */
    public Authority security(Security clzss, RequestMappingInfo info, int count) {
        String[] split = clzss.names().split(",");
        String[] urls = clzss.value().split(",");
        for (int i=0;i<urls.length;i++){
            String url = getUrl(info);
            if (url.equals(urls[i]))return null;
        }
        StringBuffer sb = new StringBuffer();
        sb.append(clzss.order()).append(count);
        return Authority.builder().pName(split[0]).pCode(String.valueOf(clzss.order())).code(sb.toString()).mOrder(0).ifMenu(false).flag(false).build();
    }

    /**
     * 普通菜单
     */
    public List<Authority> menus(Security clzss,Security security,int count){
        int order = clzss.order();
        int menu = security.menu();
        String[] split = security.names().split(",");
        List<String> index=new ArrayList<>();
        index.add(clzss.names());
        for (int i=0;i<split.length;i++) index.add(split[i]);
        List<Authority>list=new ArrayList<>();
        StringBuffer sb=new StringBuffer();
        sb.append(order);
        for (int i=2;i<=menu;i++){
            if (i!=menu){
                String pCode=sb.toString();
                sb.append("_");
                sb.append(i);
                if (security.sign()!=0) sb.append("/").append(security.sign());
                String code=sb.toString();
               list.add(Authority.builder().pName(index.get(i-2)).pCode(pCode).code(code).name(index.get(i-1)).mOrder(security.order()).ifMenu(true).flag(false).build());
            }else{
                String pCode=sb.toString();
                sb.append("_");
                sb.append(count);
                String code=sb.toString();
                list.add(Authority.builder().pName(index.get(i-2)).pCode(pCode).code(code).mOrder(security.order()).ifMenu(false).flag(false).build());
            }
        }
        return list;
    }


    /**
     * 顶级菜单
     */
    public Authority menuTop(Security security){
       return Authority
               .builder()
               .pName("")
               .pCode("")
               .code(String.valueOf(security.order()))
               .name(security.names().split(",")[0])
               .mOrder(security.order())
               .ifMenu(true)
               .flag(true)
               .build();
    }



    /**
     * 获取url
     */
    public void url(Authority a, RequestMappingInfo info){
       a.setUri(getUrl(info));
    }

    private String getUrl(RequestMappingInfo info){
        String url="";
        //获取url的Set集合，一个方法可能对应多个url
        Set<String> patterns = info.getPatternsCondition().getPatterns();
        for (String url2 : patterns)url=url2;
        return url;
    }


    /**
     * 获取ApiOperation 基本信息
     */
    public void apiOperation(Authority a, RequestMappingInfo info, Method m ){
        if (a.getIfMenu())return;
        //设置url
        url(a,info);
        ApiOperation apiOperation = m.getAnnotation(ApiOperation.class);
        Method[] me = {};
        if(apiOperation!=null) me = apiOperation.annotationType().getDeclaredMethods();
        for(Method meth : me){
            try {
                if("notes".equals(meth.getName())){
                    String color = (String) meth.invoke(apiOperation, new  Object[]{});
                    a.setDetails(color);
                }
                if("value".equals(meth.getName())){
                    String color = (String) meth.invoke(apiOperation, new  Object[]{});
                    a.setName(color);
                }
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
    }


}

