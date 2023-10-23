<script>
import svelteLogo from '../public/assets/svelte.svg'
import Counter from './lib/Counter.svelte'
import Cards from './lib/Cards.svelte';

let currentRoom = "0";
let roomOwner = true;
let cards = []
let gameStarted = false;
let players = []

function setCurrentRoom(roomId) {
    currentRoom = roomId;
}

function bin2String(array) {
  var result = "";
  for (var i = 0; i < array.length; i++) {
    result += String.fromCharCode(parseInt(array[i], 2));
  }
  return result;
}

const conn = new WebSocket("ws://192.168.15.10:8080/ws");

conn.onclose = (evt) => {
    console.log('CONNECTION CLOSED', evt)
}

conn.onmessage = (evt) => {
    console.log('MESSAGE RECEIVED', typeof evt.data);

    if (evt.data.startsWith("room")) {
        setCurrentRoom(evt.data.split(":")[1])
    } else {
        const splitData = evt.data.split(":")
      switch (splitData[0]) {
        case "BROADCAST":
            gameStarted = true;
            console.log(gameStarted)
            break;
        case "DRAWCARD":
            roomOwner = false;
            break;
        case "card":
            cards = [...cards, ...splitData[1].split("\n")]
            break
        default:
            let parsed = JSON.parse(evt.data)
            if(typeof parsed == "object") {
                players = [...parsed]
            }
            console.log(JSON.parse(evt.data))
            break;
      }
    }
}

function createRoom() {
    currentRoom = prompt("Enter room id")
    conn.send("create_room:" + currentRoom)
}

function joinRoom() {
    currentRoom = prompt("Enter room id")
    conn.send("join_room:" + currentRoom)
}

function getAllRooms() {
    conn.send("get_all_rooms")
}

function startGame() {
    conn.send("start_game")
}

function drawCard() {
    conn.send("handle_game:"+ currentRoom+":draw_card")
}
</script>

<main>
    <div>
        <a href="https://svelte.dev" target="_blank" rel="noreferrer">
            <img src={svelteLogo} class="logo svelte" alt="Svelte Logo" />
        </a>
    </div>
    <h1>Jackblack <br/>(Online Blackjack)</h1>

    <div class="card">
      {#if currentRoom == "0"}
        <Counter createRoom={createRoom} joinRoom={joinRoom} getAllRooms={getAllRooms}/>
      {:else}
        <div>
          <h2>Room ID: {currentRoom}</h2>
          <div>
              <h3>Players</h3>
                {#each players as player}
                    <p>{player}</p>
                {/each}
          </div>
        {#if roomOwner}
            <button on:click={startGame}>
                Start Game
            </button>
            <button on:click={drawCard}>
                Draw Card
            </button>
            <Cards cards={cards}/>
                {/if}
        </div>
        {/if}
</main>

<style>
.logo {
    height: 6em;
    padding: 1.5em;
    will-change: filter;
    transition: filter 700ms;
}

.logo.svelte:hover {
    filter: drop-shadow(0 0 2em #ff3e00aa);
}

button {
    margin-bottom: 2rem;
}
</style>
