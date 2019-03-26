package com.snow.golang.common.util;

import java.io.File;
import java.nio.file.Path;
import java.nio.file.Paths;

public class PathUtils {




    // 路径转换
    public static String pathToPath(String str){
        String path=null;
        if("\\".equals(File.separator)){
            //windows下
            path=str+"\\";
        }else if("/".equals(File.separator)){
            //linux下
            path=str+"/";
        }else {
            path=str;
        }
        return path;
    }

    /**
     * 获取nio path
     * @param str
     * @return
     */
    public static Path getPath(String str){
        return Paths.get(pathToPath(str));
    }



}

