<script lang="ts">
    import {telegram} from "./lib/telegram";
    import Home from "./pages/Home.svelte";
    import Navbar from "./components/Navbar.svelte";

    telegram?.ready();
    telegram?.setHeaderColor("secondary_bg_color");
    telegram?.setBackgroundColor("secondary_bg_color");
</script>

<main>
    <div class="content">
        <Home/>
    </div>
    <Navbar/>
</main>

<style>
    .content {
        flex: 1;
        overflow-y: auto;
        padding-bottom: 10px;
    }

    .content::-webkit-scrollbar {
        display: none;
    }

    main {
        height: 100vh;
        display: flex;
        flex-direction: column;
        background: var(--tg-theme-secondary-bg-color);
    }
</style>