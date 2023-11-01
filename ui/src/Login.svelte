<!-- src/Login.svelte -->
<script>
    import * as api from "./apiService";
    import * as authn from "keratin-authn";

    export let isRegistration = false; // Prop to specify whether it's a registration form

    let email = '';
    let password = '';

    async function handleSubmit() {
        // Implement your login or registration logic
        // If isRegistration is true, perform registration logic
        // Otherwise, perform login logic

        if (isRegistration) {
            await authn.signup({username:email, password});
        } else {
            await authn.login({username: email, password});
        }
    }
</script>

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