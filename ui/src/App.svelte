<!-- src/App.svelte -->
<script>
  import { onMount } from 'svelte';
  import { getTodos, addTodo, updateTodo, deleteTodo } from './apiService';
  import Login from "./Login.svelte";
  import * as authn from "keratin-authn";

  let todos = [];
  let task = '';

  authn.setHost("https://auth.todoauthn.com/");
  authn.setCookieStore("todo", {path: "/", sameSite: "None"});

  onMount(async () => {
    await authn.importSession();
    todos = await getTodos();
  });

  async function addNewTodo() {
    if (task) {
      const newTodo = await addTodo(task);
      todos = [...todos, newTodo];
      task = '';
    }
  }

  async function updateTask(id, title, description, completed) {
    await updateTodo(id, title, description, completed);
  }

  async function removeTask(id) {
    await deleteTodo(id);
    todos = todos.filter(todo => todo.id !== id);
  }
</script>
{#if authn.session() === ""}
  <Login />
{:else}
<main>
  <h1>TODO List</h1>
  <div>
    <input type="text" bind:value={task} placeholder="Enter a new task" />
    <button on:click={addNewTodo}>Add Task</button>
  </div>
  <ul>
    {#each todos as todo (todo.id)}
      <li>
        <input type="text" bind:value={todo.title} placeholder="Title" on:blur={() => updateTask(todo.id, todo.title, todo.description, todo.completed)} />
        <input type="text" bind:value={todo.description} placeholder="Description" on:blur={() => updateTask(todo.id, todo.title, todo.description, todo.completed)} />
        <input type="checkbox" bind:checked={todo.completed} on:change={() => updateTask(todo.id, todo.title, todo.description, todo.completed)} />
        {todo.title}
        <button on:click={() => removeTask(todo.id)}>Remove</button>
      </li>
    {/each}
  </ul>
</main>
{/if}