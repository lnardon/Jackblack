<script>
    import {card_images} from "../cards"
export let cards


let rooms = [{id: 1, card: "2_of_clubs"}, {id: 2, card: "5_of_spades"}]
setTimeout(() => {
    rooms = [{id: 1, card: "2_of_clubs"}, {id: 2, card: "5_of_spades"},{id: 3, card: "10_of_diamonds"}]
    setTimeout(() => {
        rooms = [{id: 1, card: "2_of_clubs"}, {id: 2, card: "5_of_spades"},{id: 3, card: "10_of_diamonds"},{id: 3, card: "3_of_spades"}]
    }, 6000)
}, 4000)
</script>

<div class="container">
    {#each cards as card}
        <div class="card_container">
            <img src={"assets/cards/"+card_images[card]+".png"} alt="Card: ">
        </div>
    {/each}
</div>

<style>
    .container{
        display: flex;
        justify-content: center;
        align-items: center;
        flex-wrap: wrap;
        gap: 2rem;
        margin-bottom: 4rem;
    }
    .card_container {
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: 2rem;
    }

    img {
        width: 200px;
        height: 300px;
        transition: all 0.5s cubic-bezier(0.075, 0.82, 0.165, 1);
        animation: spin 0.5s cubic-bezier(0.175, 0.825, 0.64, 1.275);
        cursor: pointer;
    }

    img:hover {
        transform: scale(1.05);
    }

    @keyframes spin {
        0% { 
            transform: rotateY(90deg);
            opacity: 0;
        }
        100% { 
            transform: rotateY(0deg);
            opacity: 1;
        }
    }
</style>