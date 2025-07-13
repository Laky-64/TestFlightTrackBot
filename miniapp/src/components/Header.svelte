<script>
    import {isTgUI} from "../lib/device_utils";
</script>
<div class="header" class:isTgUI>
    <slot/>
</div>

<style>
    .header {
        padding-block-start: 14px;
        padding-inline: 14px;
        width: 100%;
        display: flex;
        flex-direction: row-reverse;
        justify-content: space-between;
    }

    .header.isTgUI {
        background: var(--tg-theme-section-bg-color);
    }
</style>