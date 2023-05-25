// Функция для отправки запроса на сервер и обработки ответа
async function sendRequest(url, method, data) {
    try {
      const response = await fetch(url, {
        method: method,
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
      });
  
      if (!response.ok) {
        throw new Error('Ошибка сервера');
      }
  
      const responseData = await response.json();
      return responseData;
    } catch (error) {
      console.error(error);
      throw error;
    }
  }
  
  // Функция для создания новой задачи
  async function createTask() {
    const title = document.getElementById('task-title').value;
    const description = document.getElementById('task-description').value;
  
    const data = {
      title: title,
      description: description
    };
  
    try {
      await sendRequest('/tasks', 'POST', data);
      // Обновление списка задач
      await fetchTasks();
    } catch (error) {
      console.error(error);
    }
  }
  
  // Функция для получения списка задач
  async function fetchTasks() {
    try {
      const tasksContainer = document.getElementById('task-list');
      tasksContainer.innerHTML = ''; // Очистка контейнера
  
      const tasks = await sendRequest('/tasks', 'GET', null);
  
      tasks.forEach((task) => {
        const taskItem = document.createElement('div');
        taskItem.classList.add('task-item');
        taskItem.innerHTML = `
          <h3>${task.title}</h3>
          <p>${task.description}</p>
          <button onclick="deleteTask(${task.id})">Удалить</button>
        `;
        tasksContainer.appendChild(taskItem);
      });
    } catch (error) {
      console.error(error);
    }
  }
  
  // Функция для удаления задачи по идентификатору
  async function deleteTask(taskId) {
    try {
      await sendRequest(`/tasks/${taskId}`, 'DELETE', null);
      // Обновление списка задач
      await fetchTasks();
    } catch (error) {
      console.error(error);
    }
  }
  
  // Обработчик события отправки формы для создания задачи
  document.getElementById('task-form').addEventListener('submit', function(event) {
    event.preventDefault();
    createTask();
  });
  
  // Загрузка списка задач при загрузке страницы
  fetchTasks();
  