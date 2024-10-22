<script>
    import Drawer, {
        AppContent,
        Content,
        Header,
        Title,
        Subtitle,
    } from "@smui/drawer";
    import Button, { Label }from "@smui/button";
    import List, { Item, Text } from '@smui/list';

    let open = true;
    let active = 'General';
    const channels = ['General', 'Random', 'Dev Chat']

    let loggedInUser = {
        username: 'johndoe',
        displayName: 'John Doe',
        role: 'user',
    };

    function setActive(value) {
        active = value;
    }
  // Log to verify if the state-related logic is working
  console.log('Is Drawer Open?', open);
  console.log('Active Channel:', active);
</script>

<!-- Layout Structure -->
 <div class="drawer-container">

    <!-- Collabsible Sidebar -->
    <Drawer variant="dismissable" bind:open={open}>
        <!-- Sidebar Header -->
        <Header>
            <Title>Chat App</Title>
            <Subtitle>Friends and Channels</Subtitle>
        </Header>

        <Content>
            <List>
                {#each channels as channel}
                    <Item
                        href="javascript:void(0)"
                        on:click={() => setActive(channel)}
                        activated={active === channel}
                    >
                        <Text>{channel}</Text>
                    </Item>
                {/each}
            </List>
        </Content>
    </Drawer>

    <AppContent class="app-content">
        <main class="main-content">
            <!-- Toggle button to open/close the drawer -->
            <Button on:click={() => (open = !open)}>
                <Label>Toggle Sidebar</Label>
            </Button>
            
            <!-- Show active channel -->
            <pre class="status">Active Channel: {active}</pre>
            <slot />
        </main>
    </AppContent>
 </div>
  
<style>
/* Import global styles from Svelte Material UI */
@import '@smui/button/styles.css';
@import '@smui/textfield/styles.css';
@import '@smui/card/styles.css';
@import '@smui/formfield/styles.css';
@import '@smui/drawer/styles.css';
@import '@smui/list/styles.css';


  /* Overall container for the Drawer and main content */
  .drawer-container {
    position: relative;
    display: flex;
    min-height: 100vh;  /* Full viewport height */
    overflow: hidden;
    z-index: 0;
  }

  /* Applies Material UI styling for the main content area */
  * :global(.app-content) {
    flex: auto;
    overflow: auto;
    position: relative;
    flex-grow: 1;
  }

  /* Main content area between sidebar and content */
  .main-content {
    overflow: auto;
    padding: 16px;
    height: 100%;
    box-sizing: border-box;
  }
  
  /* Inline styles for status output */
  .status {
    margin-top: 1rem;
  }
</style>
  
  <!-- Slot renders the content of the routed pages
  <slot /> -->
