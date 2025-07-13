<script lang="ts">
    import {isTgUI} from "../lib/device_utils";
</script>

<div class="navbar" class:isTgUI>

</div>

<style>
    .navbar {
        height: 60px;
        width: 100%;
        background: var(--tg-theme-bottom-bar-bg-color);
        border-top: 1px solid var(--tg-theme-section-separator-color);
    }

    .navbar.isTgUI {
        background: var(--tg-theme-section-bg-color);
    }
</style>