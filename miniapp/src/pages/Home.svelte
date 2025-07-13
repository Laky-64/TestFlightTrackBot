<script lang="ts">
    import {currentUser, telegram, type PopupParams} from "../lib/telegram";
    import UserIcon from '../components/UserIcon.svelte'
    import Header from "../components/Header.svelte";
    import StickerContainer from "../components/StickerContainer.svelte";
    import ListView from "../components/ListView.svelte";
    import Divider from "../components/Divider.svelte";
    import Button from "../components/Button.svelte";
    import ItemView from "../components/ItemView.svelte";
    import {isTgUI} from "../lib/device_utils";

    let deletable = $state(isTgUI);
    let textButton = $state('Edit');

    let items = $state([
        {
            id: "1",
            title: "Telegram",
            tID: "u6iogfd0",
            highlight: true,
            desc: "Available",
            icon: "https://is1-ssl.mzstatic.com/image/thumb/Purple211/v4/0e/5e/d2/0e5ed270-208b-5f0e-4f43-e0c51a598f7a/AppIconLLC-0-0-1x_U007emarketing-0-8-0-85-220.png/152x152bb-80.png"
        },
        {
            id: "2",
            title: "WhatsApp",
            tID: "oscYikr0",
            desc: "Last availability 1 week ago",
            icon: "https://is1-ssl.mzstatic.com/image/thumb/Purple221/v4/24/d6/69/24d66960-dbf7-096c-886f-14e023628808/AppIconSMB-0-0-1x_U007ephone-0-0-0-1-0-0-sRGB-0-85-220.png/152x152bb-80.png"
        }
    ]);

    function startDeletion() {
        deletable = !deletable;
        textButton = deletable ? 'Done' : 'Edit';
    }

    function removeItem(id: string) {
        telegram.showPopup(
            {
                title: "Delete Link",
                message: "Are you sure you want to delete this link?",
                buttons: [
                    {
                        id: "delete",
                        text: "Delete",
                        type: "destructive",
                    },
                    {
                        type: "cancel",
                    }
                ]
            },
            (result: string) => {
                if (result !== "delete") return;
                items = items.filter(item => item.id !== id);
            },
        );
    }
</script>
<Header>
    <Button type="opaque"><UserIcon user={currentUser} /></Button>
    {#if !isTgUI}
        <Button bind:text={textButton} click={() => startDeletion()}/>
    {/if}
</Header>
<StickerContainer
    title="Track Your Favorite Betas "
    desc="Get instant updates when TestFlight links open, close, or fill up stay ahead in testing!"
    sticker="DuckHome"
/>
<Divider/>
<ListView header="Links">
    <ItemView bind:deletable icon="add" title="Add More"/>
    {#each items as item (item.id)}
        <ItemView
            {...item}
            bind:deletable
            on_delete={() => removeItem(item.id)}
        />
    {/each}
</ListView>