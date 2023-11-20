<!-- src/App.svelte -->
<script>
  import { onMount } from 'svelte';
  import { getTodos, addTodo, updateTodo, deleteTodo } from './apiService';
  import * as authn from "keratin-authn";

  let todos = [];
  let task = '';

  authn.setHost("https://auth.todoauthn.com/");
  authn.setCookieStore("todo", {path: "/", sameSite: "None", domain: "todoauthn.com"});

  onMount(async () => {
    await authn.importSession();
    todos = await getTodos() ?? [];
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

  export let isRegistration = false; // Prop to specify whether it's a registration form

  let email = '';
  let password = '';

  async function handleSubmit() {
    if (isRegistration) {
      await authn.signup({username:email, password:password});
    } else {
      await authn.login({username: email, password:password});
    }
    window.location.reload();
  }

  async function logout() {
    await authn.logout();
    window.location.reload();
  }
</script>
<!-- TODO: this is not updated after login until we refresh the page a couple times -->
{#if authn.session() === ""}
  <div>
    <h2>{isRegistration ? 'Register' : 'Login'}</h2>
    <form>
      <label for="email">Email:</label>
      <input type="email" id="email" bind:value={email} />

      <label for="password">Password:</label>
      <input type="password" id="password" bind:value={password} />

      <button on:click={handleSubmit}>
        {isRegistration ? 'Register' : 'Login'}
      </button>


      <label>
        <input type="checkbox" bind:checked={isRegistration} />
        I'm new here
      </label>
    </form>
  </div>
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
  <div>
    <button on:click={logout}>Logout</button>
  </div>
</main>
{/if}