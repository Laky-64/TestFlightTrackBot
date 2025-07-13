<script lang="ts">
    import {isTgUI} from "../lib/device_utils";

    const { header, children } : {header: string, children: any} = $props();
</script>

<div class="listView" class:isTgUI>
    <p>{isTgUI ? header : header.toUpperCase()}</p>
    <div>
        {@render children?.()}
    </div>
</div>

<style>
    .listView {
        padding-inline: 14px;
        padding-block: 10px;
    }

    .listView.isTgUI {
        padding-inline: 0;
        background: var(--tg-theme-section-bg-color);
    }

    .listView.isTgUI > p {
        font-family: var(--font-bold), serif;
        font-size: 15px;
        margin-top: 10px;
    }

    .listView > p {
        color: var(--tg-theme-section-header-text-color);
        margin: 0 0 8px 15px;
        font-size: 14px;
    }

    .listView > div {
        background: var(--tg-theme-section-bg-color);
        width: 100%;
        border-radius: 14px;
        overflow: hidden;
    }

    .listView.isTgUI > div {
        border-radius: 0;
    }
</style>