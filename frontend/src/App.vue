<template>
  <main style="padding: 1rem; font-family: Arial, sans-serif;">
    <h1>TodoMVC</h1>

    <form @submit.prevent="addTodo" style="margin-bottom:1rem;">
      <input
        v-model="newText"
        placeholder="What needs to be done?"
        required
        style="padding:0.5rem; width: 250px;"
      />
      <input
        v-model="newDue"
        type="date"
        style="padding:0.5rem; margin-left:0.5rem;"
      />
      <button type="submit" style="padding:0.5rem;">Add</button>
    </form>

    <ul>
      <li v-for="t in todos" :key="t.id">
        <label>
          <input type="checkbox" v-model="t.done" @change="updateTodo(t)" />
          <span :style="{ textDecoration: t.done ? 'line-through' : 'none' }">
            {{ t.text }} (due: {{ t.due_date }})
          </span>
        </label>
        <button @click="deleteTodo(t.id)" style="margin-left:0.5rem;">×</button>
      </li>
    </ul>
  </main>
</template>

<script setup>
import { ref, onMounted } from 'vue';

const API_BASE = "/api";
const todos = ref([]);
const newText = ref('');
const newDue = ref('');

const fetchTodos = async () => {
  const res = await fetch(`${API_BASE}/todos`);
  todos.value = await res.json();
};

const addTodo = async () => {
  const res = await fetch(`${API_BASE}/todos`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ text: newText.value, due_date: newDue.value }),
  });
  if (res.ok) {
    newText.value = '';
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
ul {
  list-style: none;
  padding: 0;
}
li {
  margin: 0.5rem 0;
}
button {
  cursor: pointer;
}
</style>
