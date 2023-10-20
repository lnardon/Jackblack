<script>
  import svelteLogo from '../public/assets/svelte.svg'
  import Counter from './lib/Counter.svelte'
  import Cards from './lib/Cards.svelte';

  let currentRoom = 0;

  function setCurrentRoom(roomId){
    currentRoom = roomId;
  }
</script>

<main>
  <div>
    <a href="https://svelte.dev" target="_blank" rel="noreferrer">
      <img src={svelteLogo} class="logo svelte" alt="Svelte Logo" />
    </a>
  </div>
  <h1>Jackblack (Online Blackjack)</h1>

  <div class="card">
    {#if currentRoom == 0}
    <Counter setCurrentRoom={setCurrentRoom}/>
    {:else}
    <Cards />
    {/if}
  </div>

</main>

<style>
  .logo {
    height: 6em;
    padding: 1.5em;
    will-change: filter;
    transition: filter 300ms;
  }
  .logo:hover {
    filter: drop-shadow(0 0 2em #646cffaa);
  }
  .logo.svelte:hover {
    filter: drop-shadow(0 0 2em #ff3e00aa);
  }
</style>
