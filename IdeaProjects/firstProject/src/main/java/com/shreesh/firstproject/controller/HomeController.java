package com.shreesh.firstproject.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/first")
public class HomeController {

    @GetMapping("/project")
    public String home(){

        return "Dont worry about the results just keep hustling........";
    }
}