---
title: 【study】Spring学习day01
date: 2024-01-28 10:57:22
tags:
  - Study
  - Java
  - Spring框架
---

# 【study】Spring学习day01

- 传统JavaWeb开发的困惑
- IoC、DI和AOP思想提出
- Spring框架诞生

## 传统JavaWeb开发的困惑以及解决方案

``` java
// 用户账户信息修改业务方法
public void updateUserInfo(User user) {
    try {
        // 开启事务
        DaoUtils.openTransaction();
        // 获得UserDao执行插入User数据到数据库操作
        UserDao userDao = new UserDaoImpl();
        userDao.updateUserInfo(user);
        // 修改成功后, 向用户行为日志表中插入一条数据, 内容: 修改时间等信息
        UserLog userLog = new UserLogImpl();
        UserLog.recodeUserUpdate(user);
        // 提交事务
        DaoUtils.commit();
    } catch(Exception e) {
        // 回滚事务
        DaoUtils.rollback();
        // 向异常日志表中插入数据
        ExceptionLog exceptionLog = new ExceptionLogImpl();
        exceptionLog.recodeException(this, e);
    }
}
```

``` java
// 用户注册业务方法
public void regist(User user) {
    try {
        // 开启事务
        DaoUtils.openTransaction();
        // 获得UserDao执行插入User数据到数据库操作
        UserDao userDao = new UserDaoImpl();
        userDao.addUser(user);
        // 修改成功后, 向用户行为日志表中插入一条数据, 内容: 时间, 用户, 注册行为
        UserLog userLog = new UserLogImpl();
        UserLog.recodeUserRegist(user);
        // 注册成功后, 向用户邮箱发送一封激活邮件
        CommonUtils.sendEmail(user);
        // 提交事务
        DaoUtils.commit();
    } catch(Exception e) {
        // 回滚事务
        DaoUtils.rollback();
        // 向异常日志表中插入数据
        ExceptionLog exceptionLog = new ExceptionLogImpl();
        exceptionLog.recodeException(this, e);
    }
}
```

代码是两个业务层的代码, 主业务均为第7\~8行

现在存在一些问题: 以`updateUserInfo()`方法举例

``` java
UserDao userDao = new UserDaoImpl();
UserLog userLog = new UserLogImpl();
ExceptionLog exceptionLog = new ExceptionLogImpl();
```

这里面耦合度很高:

- 层与层之间紧密耦合在了一起, 接口与具体实现紧密耦合在一起

    > 解决思路: 程序代码中不要动手new对象，第三方根据要求为程序提供需要的Bean对象

- 通用的事务功能耦合在业务代码中，通用的日志功能耦合在业务代码中

    > 程序代码中不要动手new对象，第三方根据要求为程序提供需要的Bean对象的代理对象

### IoC思想

Inversion of Control，控制反转，强调的是原来在程序中创建Bean的权利反转给第三方

### DI思想

Dependency Injection，依赖注入，强调的Bean之间关系，这种关系第三方负责去设置

### AOP思想

Aspect Oriented Programming，面向切面编程，功能的横向抽取，主要的实现方式是Proxy

## 基于XML管理Bean

IDE中新建Java项目，并且新建子模块`day01`, 父模块中的`pom.xml`引入一些依赖

==jdk版本: 21.0.2==

``` xml
<dependencies>
    <dependency>
        <groupId>org.springframework</groupId>
        <artifactId>spring-context</artifactId>
        <version>6.1.2</version>
    </dependency>

    <dependency>
        <groupId>org.junit.jupiter</groupId>
        <artifactId>junit-jupiter-api</artifactId>
        <version>5.8.2</version>
    </dependency>
</dependencies>
```

建立一些简单的文件

在`day01.src.main.java.org.example`文件夹下

``` java
// service.UserService.Java
package org.example.service;

public interface UserService {
    void add();
}

```

``` java
// service.impl.UserSerivceImpl
package org.example.service.impl;

import org.example.service.UserService;

public class UserServiceImpl implements UserService {
	public void add() {
        System.out.println("add......");
    }
}

```

`day01/src/main/resources/`文件夹内添加xml配置文件, 文件名没有要求，这里叫做`beans.xml`

``` xml
<?xml version="1.0" encoding="UTF-8"?>
<beans xmlns="http://www.springframework.org/schema/beans"
       xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
       xsi:schemaLocation="http://www.springframework.org/schema/beans http://www.springframework.org/schema/beans/spring-beans.xsd">
    <bean class="org.example.service.impl.UserServiceImpl" id="userService">

    </bean>
</beans>
```

`bean`标签中，class指定要创建对象的类，id是唯一标识



``` java
// test.TestUserService
public class TestUserService {
    @Test
    public void testUserService() {
        // 加载Spring配置文件, 对象创建
        ApplicationContext context = new ClassPathXmlApplicationContext("beans.xml");

        // 获取创建的对象
        // context.getBean("userService")的返回值为Object类型, 需要强制类型转换
        UserService userService = (UserService) context.getBean("userService");
        System.out.println(userService);

        // 使用对象调用方法进行调试
        userService.add();
    }

}

```

