package com.sonifoy.ranking;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class RankingServiceApplication {

    public static void main(String[] args) {
        System.out.println("--- Environment Variable Check ---");
        System.out.println("DATABASE_URL: " + (System.getenv("DATABASE_URL") != null));
        System.out.println("EUREKA_CLIENT_SERVICEURL_DEFAULTZONE: " + System.getenv("EUREKA_CLIENT_SERVICEURL_DEFAULTZONE"));
        System.out.println("SPRING_PROFILES_ACTIVE: " + System.getenv("SPRING_PROFILES_ACTIVE"));
        System.out.println("---------------------------------");
        
        SpringApplication.run(RankingServiceApplication.class, args);
    }

}
