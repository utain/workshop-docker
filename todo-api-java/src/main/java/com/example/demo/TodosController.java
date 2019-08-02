package com.example.demo;

import java.util.ArrayList;
import java.util.List;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RequestMapping("/api")
@RestController
public class TodosController {
	class Todo {
		private String title;
		private Boolean done;

		public Todo(String title, Boolean done) {
			this.title = title;
			this.done = done;
		}

		public String getTitle() {
			return title;
		}

		public void setTitle(String title) {
			this.title = title;
		}

		public Boolean getDone() {
			return done;
		}

		public void setDone(Boolean done) {
			this.done = done;
		}

	}

	@GetMapping("/tasks")
	public List<Todo> getTodo() {
		List<Todo> todos = new ArrayList<>();
		todos.add(new Todo("Learn SpringBoot", true));
		todos.add(new Todo("Learn ReactJs", true));
		todos.add(new Todo("Learn Angular", true));
		todos.add(new Todo("Learn Go!s", false));
		return todos;
	}

}
