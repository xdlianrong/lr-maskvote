package com.xdlr.xdublockchainlab.maskvote.config;

import org.springframework.web.servlet.HandlerInterceptor;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.servlet.http.HttpSession;

public class LoginInterceptor implements HandlerInterceptor {


    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) throws Exception {

        HttpSession session = request.getSession();
        Object loginUser = session.getAttribute("loginUser");//这里是存在问题的，但是现在我着急过代码，所以一会再来解决掉这个问题
        //但是这个问题是具有启发性质的
        //它可能会关联到整个现在的这个项目
        if (loginUser != null){
            return true;
        }
        request.setAttribute("msg", "请先登录！");
        request.getRequestDispatcher("/").forward(request, response);
        return false;
    }
}
