package com.impacta.todo.services;

import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.Arrays;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.impacta.todo.domain.Todo;
import com.impacta.todo.repositories.TodoRepository;

@Service
public class DBService {

	@Autowired
	private TodoRepository todoRepository;

	public void instanciaBaseDeDados() {

		DateTimeFormatter formatter = DateTimeFormatter.ofPattern("dd/MM/yyyy HH:mm");

		Todo t1 = new Todo(0, "Estudar", "Estudar SpringBoot 2.x",
				LocalDateTime.parse("25/03/2022 11:40", formatter), false);
		Todo t2 = new Todo(0, "Ler", "Ler livro de desenvolvimento pessoal",
				LocalDateTime.parse("12/05/2022 11:40", formatter), true);
		Todo t3 = new Todo(0, "Execicios","Praticar exercicios fisicos",
				LocalDateTime.parse("13/05/2022 11:40", formatter), false);
		Todo t4 = new Todo(0, "Meditar", "Meditar durante 30 minutos pela manha",
				LocalDateTime.parse("14/05/2022 11:40", formatter), true);
		
		todoRepository.saveAll(Arrays.asList(t1, t2, t3, t4));
	}

}
