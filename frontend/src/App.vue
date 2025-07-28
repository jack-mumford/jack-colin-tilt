<template>
  <div class="app-container">
    <main class="main-content">
      <h1 class="app-title">TodoMVC</h1>

      <form @submit.prevent="addTodo" class="todo-form">
        <div class="input-group">
          <input
            v-model="newText"
            placeholder="What needs to be done?"
            required
            class="input-field input-text"
          />
          <input
            v-model="newDescription"
            placeholder="Description (optional)"
            class="input-field input-description"
          />
          <input
            v-model="newDue"
            type="date"
            class="input-field input-date"
          />
          <button type="submit" class="add-button">Add Todo</button>
        </div>
      </form>

      <div class="todos-container">
        <div v-for="t in todos" :key="t.id" class="todo-card">
          <div class="todo-content">
            <label class="todo-label">
              <input 
                type="checkbox" 
                v-model="t.done" 
                @change="updateTodo(t)" 
                class="todo-checkbox"
              />
              <div class="todo-text">
                <span 
                  class="todo-title" 
                  :class="{ 'completed': t.done }"
                >
                  {{ t.text }}
                </span>
                <span v-if="t.description" class="todo-description">
                  {{ t.description }}
                </span>
                <span class="todo-due">Due: {{ t.due_date }}</span>
              </div>
            </label>
          </div>
          <button @click="deleteTodo(t.id)" class="delete-button">
            ×
          </button>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';

const API_BASE = "/api";
const todos = ref([]);
const newText = ref('');
const newDescription = ref('');
const newDue = ref('');

const fetchTodos = async () => {
  const res = await fetch(`${API_BASE}/todos`);
  todos.value = await res.json();
};

const addTodo = async () => {
  const res = await fetch(`${API_BASE}/todos`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ text: newText.value, description: newDescription.value, due_date: newDue.value }),
  });
  if (res.ok) {
    newText.value = '';
    newDescription.value = '';
    newDue.value = '';
    await fetchTodos();
  }
};

const updateTodo = async (todo) => {
  await fetch(`${API_BASE}/todos/${todo.id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(todo),
  });
};

const deleteTodo = async (id) => {
  await fetch(`${API_BASE}/todos/${id}`, { method: 'DELETE' });
  await fetchTodos();
};

onMounted(fetchTodos);
</script>

<style scoped>
.app-container {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  justify-content: center;
  align-items: flex-start;
  padding: 2rem 1rem;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

.main-content {
  background: white;
  border-radius: 12px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
  padding: 2rem;
  width: 100%;
  max-width: 600px;
  margin-top: 2rem;
}

.app-title {
  text-align: center;
  color: #333;
  margin-bottom: 2rem;
  font-size: 2.5rem;
  font-weight: 300;
  letter-spacing: -1px;
}

.todo-form {
  margin-bottom: 2rem;
}

.input-group {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.input-field {
  padding: 0.75rem 1rem;
  border: 2px solid #e1e5e9;
  border-radius: 8px;
  font-size: 1rem;
  transition: border-color 0.3s ease, box-shadow 0.3s ease;
  outline: none;
}

.input-field:focus {
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.input-text {
  font-weight: 500;
}

.input-description {
  color: #666;
}

.input-date {
  color: #666;
}

.add-button {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  padding: 0.75rem 2rem;
  border-radius: 8px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.add-button:hover {
  transform: translateY(-2px);
  box-shadow: 0 5px 15px rgba(102, 126, 234, 0.4);
}

.add-button:active {
  transform: translateY(0);
}

.todos-container {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.todo-card {
  background: #f8f9fa;
  border: 1px solid #e9ecef;
  border-radius: 10px;
  padding: 1.25rem;
  display: flex;
  align-items: center;
  justify-content: space-between;
  transition: all 0.3s ease;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
}

.todo-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  border-color: #667eea;
}

.todo-content {
  flex: 1;
}

.todo-label {
  display: flex;
  align-items: flex-start;
  gap: 1rem;
  cursor: pointer;
}

.todo-checkbox {
  width: 20px;
  height: 20px;
  margin-top: 2px;
  cursor: pointer;
  accent-color: #667eea;
}

.todo-text {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.todo-title {
  font-size: 1.1rem;
  font-weight: 500;
  color: #333;
  transition: all 0.3s ease;
}

.todo-title.completed {
  text-decoration: line-through;
  color: #999;
}

.todo-description {
  font-size: 0.9rem;
  color: #666;
  font-style: italic;
}

.todo-due {
  font-size: 0.85rem;
  color: #888;
  font-weight: 500;
}

.delete-button {
  background: #ff4757;
  color: white;
  border: none;
  border-radius: 50%;
  width: 32px;
  height: 32px;
  font-size: 1.2rem;
  font-weight: bold;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s ease;
  margin-left: 1rem;
}

.delete-button:hover {
  background: #ff3742;
  transform: scale(1.1);
}

.delete-button:active {
  transform: scale(0.95);
}

@media (min-width: 768px) {
  .input-group {
    flex-direction: row;
    align-items: center;
  }
  
  .input-text {
    flex: 2;
  }
  
  .input-description {
    flex: 1.5;
  }
  
  .input-date {
    flex: 1;
  }
  
  .add-button {
    flex: 0 0 auto;
    white-space: nowrap;
  }
}
</style>
