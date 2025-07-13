<script lang="ts">
    import { slide } from 'svelte/transition';
    import Button from "./Button.svelte";
    import {isTgUI, isDesktop} from "../lib/device_utils";

    let {
        icon,
        title,
        desc,
        tID,
        on_click,
        on_delete,
        deletable = $bindable(),
        highlight = false,
    } : {
        icon: string,
        title?: string,
        desc?: string,
        tID?: string,
        on_click?: any,
        on_delete?: any,
        deletable?: boolean,
        highlight?: boolean,
    } = $props();

    function onKeyDelete(e: KeyboardEvent) {
        if ((e.key === "Enter" || e.key === " ") && on_delete) {
            e.preventDefault();
            on_delete();
        }
    }
</script>

<div class="itemView" class:deletable class:isTgUI out:slide={{ duration: 250 }}>
    <Button click={() => {if (on_click) on_click()}}>
        {#if icon === "add"}
            <svg class="icon" width="20px" height="20px" viewBox="0 0 20 20" fill="var(--tg-theme-accent-text-color)" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd" clip-rule="evenodd" d="M10 20C9.44771 20 9 19.5523 9 19L9 11H1C0.447716 11 0 10.5523 0 10C0 9.44771 0.447716 9 1 9H9V1C9 0.447716 9.44771 0 10 0C10.5523 0 11 0.447716 11 1V9L19 9C19.5523 9 20 9.44771 20 10C20 10.5523 19.5523 11 19 11L11 11L11 19C11 19.5523 10.5523 20 10 20Z"/>
            </svg>
            <div>
                <p class="addMoreText">{title}</p>
            </div>
        {:else}
            {#if !isTgUI}
                <div class="deleteBtn" role="button" tabindex="0" onkeydown={onKeyDelete} onclick={() => {if (on_delete) on_delete()}} class:deletable class:isTgUI>
                    <div></div>
                </div>
            {/if}
            <img src="{icon}" alt="ItemIcon">
            <div>
                <p class="itemTitle">{title} [<span>#{tID}</span>]</p>
                <p class="itemDesc" class:highlight>{desc}</p>
            </div>
            {#if isTgUI && deletable}
                <Button click={() => {if (on_delete) on_delete()}} type="default-rounded">
                    <svg class="deleteBtn" class:isTgUI class:isDesktop width="18" height="22" viewBox="0 0 18 22" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path fill-rule="evenodd" clip-rule="evenodd" d="M5.19995 3.19999C5.19995 1.6536 6.45355 0.399994 7.99995 0.399994H9.99995C11.5463 0.399994 12.8 1.6536 12.8 3.19999V3.39999H17C17.4418 3.39999 17.7999 3.75817 17.7999 4.19999C17.7999 4.64182 17.4418 4.99999 17 4.99999H0.999951C0.558124 4.99999 0.199951 4.64182 0.199951 4.19999C0.199951 3.75817 0.558124 3.39999 0.999951 3.39999H5.19995V3.19999ZM11.2 3.19999V3.39999H6.79995V3.19999C6.79995 2.53725 7.33721 1.99999 7.99995 1.99999H9.99995C10.6627 1.99999 11.2 2.53725 11.2 3.19999ZM2.79544 7.11466C2.74831 6.67535 2.35397 6.35743 1.91467 6.40456C1.47536 6.45169 1.15744 6.84602 1.20456 7.28533L2.51481 19.4987C2.66743 20.9213 3.86805 22 5.29883 22H12.7168C14.1489 22 15.3502 20.9193 15.5012 19.4951L16.7955 7.28432C16.8421 6.84496 16.5237 6.45102 16.0843 6.40445C15.645 6.35788 15.251 6.6763 15.2045 7.11566L13.9101 19.3265C13.8454 19.9368 13.3305 20.4 12.7168 20.4H5.29883C4.68564 20.4 4.17109 19.9377 4.10568 19.328L2.79544 7.11466Z"/>
                    </svg>
                </Button>
            {/if}
        {/if}
    </Button>
</div>

<style>
    /** Item View Button **/
    .itemView {
        position: relative;
        width: 100%;
        --delete-width: 45px;
    }

    .itemView.deletable > :global(div.default:not(.isTgUI):not(:has(.icon)):active > div) {
        transform: scale(1);
    }

    .itemView.isTgUI > :global(div:hover:has(div > div:last-child:hover)::after),
    .itemView.deletable > :global(div:has(div.deleteBtn):active::after) {
        opacity: 0;
    }

    /** Divider line **/
    .itemView:not(:last-child)::after {
        content: "";
        position: absolute;
        bottom: 0;
        left: 68px;
        height: 1px;
        right: 0;
        transition: transform 250ms ease;
        background: var(--tg-theme-section-separator-color);
    }

    .itemView:not(:has(.icon)):not(.isTgUI).deletable:not(:last-child)::after {
        transform: translateX(var(--delete-width));
    }

    /** Delete button **/
    .deleteBtn.isTgUI {
        margin-inline: 7px;
        margin-block: 5px;
        fill: var(--tg-theme-destructive-text-color);
    }

    .itemView > :global(div > div > div:last-child:has(.deleteBtn)) {
        margin-left: auto;
        margin-right: 24px;
        padding: 5px;
    }

    .itemView > :global(div > div > div.isDesktop:last-child:hover > div > svg) {
        fill: var(--tg-theme-text-color);
    }

    .deleteBtn.isDesktop {
        fill: var(--tg-theme-hint-color);
        transition: fill 120ms ease-in-out;
    }

    .deleteBtn:not(.isTgUI):not(:first-child) {
        display: none;
    }

    :not(svg).deleteBtn:not(.isTgUI) {
        width: 0;
        display: flex;
        align-items: center;
        justify-content: center;
        height: var(--delete-width);
        position: relative;
        transition: width 250ms ease;
    }

    :not(svg).deleteBtn:not(.isTgUI) > div {
        position: absolute;
        background: var(--tg-theme-destructive-text-color);
        width: 24px;
        height: 24px;
        border-radius: 50%;
        margin-left: 14px;
        transform: translateX(calc(-14px - var(--delete-width)));
        transition: transform 210ms ease 0ms;
    }

    :not(svg).deleteBtn:not(.isTgUI) >:after {
        content: "";
        position: absolute;
        width: 12px;
        height: 2px;
        left: calc(50% - 6px);
        top: calc(50% - 1px);
        border-radius: 1px;
        background: var(--tg-theme-text-color);
    }

    .deleteBtn:not(.isTgUI).deletable {
        width: var(--delete-width);
    }

    .deleteBtn:not(.isTgUI).deletable > div {
        transform: translateX(0);
        transition: transform 400ms ease;
    }

    /** Icons and Description **/
    .icon {
        margin-block: 15px;
        margin-inline: 24px;
    }

    img {
        height: 40px;
        border-radius: 50%;
        margin-inline: 14px;
        margin-block: 13px;
    }

    p:first-child {
        margin: 0;
        color: var(--tg-theme-text-color);
        font-size: 17px;
        font-family: var(--font-medium), serif;
    }

    p:first-child > span {
        color: var(--tg-theme-subtitle-text-color);
    }

    p:first-child.addMoreText {
        color: var(--tg-theme-accent-text-color);
    }

    .itemDesc {
        margin: 2px 0 0;
        font-size: 15px;
        color: var(--tg-theme-subtitle-text-color);
    }

    .itemDesc.highlight {
        color: var(--tg-theme-accent-text-color);
    }
</style>