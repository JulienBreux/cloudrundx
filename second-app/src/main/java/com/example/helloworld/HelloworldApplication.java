package com.example.helloworld;

import org.springframework.beans.factory.annotation.Value; // +
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.GetMapping; // +
import org.springframework.web.bind.annotation.RestController; // +

@SpringBootApplication
public class HelloworldApplication {

	// START
	@Value("${NAME:World}")
	String name;

	@RestController
	class HelloworldController {
		@GetMapping("/")
		String hello() {
			return "Hello " + name + " from second app!";
		}
	}
	// END

	public static void main(String[] args) {
		SpringApplication.run(HelloworldApplication.class, args);
	}

}
