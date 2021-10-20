import Link from 'next/link'
import * as React from 'react';
import PropTypes from 'prop-types';
import {
    AppBar,
    Box,
    Toolbar,
    IconButton,
    Typography,
    Divider,
    List,
    Drawer,
    ListItem,
    ListItemIcon,
    ListItemText
} from "@mui/material"

const drawerWidth = 240;

function Sidebar(props) {
    const { window } = props;
    const [mobileOpen, setMobileOpen] = React.useState(false);

    const handleDrawerToggle = () => {
        setMobileOpen(!mobileOpen);
    };

    const drawer = (
        <div>
            <Toolbar />
            <div className={"brand-logo"}><img src="/logo.png" alt="matrixforce" />
                <h3>Matrixforce 2.0</h3></div>

            <Divider />
            <List>
                {['Dashboard', 'Inbox', 'CRM', 'Teams' , 'Channels' , 'Boardcast' , "Bot-Builder"].map((text, index) => (
                    <ListItem button key={text}>
                        <ListItemIcon>
                            {/*{index % 2 === 0 ? <InboxIcon /> : <MailIcon />}*/}
                        </ListItemIcon>
                        <ListItemText primary={text} />
                    </ListItem>
                ))}
            </List>
            {/*<Divider />*/}
            {/*<List>*/}
            {/*    {['All mail', 'Trash', 'Spam'].map((text, index) => (*/}
            {/*        <ListItem button key={text}>*/}
            {/*            <ListItemIcon>*/}
            {/*                /!*{index % 2 === 0 ? <InboxIcon /> : <MailIcon />}*!/*/}
            {/*            </ListItemIcon>*/}
            {/*            <ListItemText primary={text} />*/}
            {/*        </ListItem>*/}
            {/*    ))}*/}
            {/*</List>*/}
        </div>
    );

    const container = window !== undefined ? () => window().document.body : undefined;

    return (
        <Box sx={{ display: 'flex' }}>
            {/*<CssBaseline />*/}
            <AppBar
                position="fixed"
                sx={{
                    width: { sm: `calc(100% - ${drawerWidth}px)` },
                    ml: { sm: `${drawerWidth}px` },
                    background:{md:"white"},
                    color:{sm:"black"}
                }}
            >
                <Toolbar>
                    <IconButton
                        color="inherit"
                        aria-label="open drawer"
                        edge="start"
                        onClick={handleDrawerToggle}
                        sx={{ mr: 2, display: { sm: 'none' } }}
                    >
                        {/*<MenuIcon />*/}
                    </IconButton>
                    <h3>
                        MatrixForce
                    </h3>
                </Toolbar>
            </AppBar>
            <Box
                component="nav"
                sx={{ width: { sm: drawerWidth }, flexShrink: { sm: 0 } }}
                aria-label="mailbox folders"
            >
                {/* The implementation can be swapped with js to avoid SEO duplication of links. */}
                <Drawer
                    container={container}
                    variant="temporary"
                    open={mobileOpen}
                    onClose={handleDrawerToggle}
                    ModalProps={{
                        keepMounted: true, // Better open performance on mobile.
                    }}
                    sx={{
                        display: { xs: 'block', sm: 'none' },
                        '& .MuiDrawer-paper': { boxSizing: 'border-box', width: drawerWidth },
                    }}
                >
                    {drawer}
                </Drawer>
                <Drawer
                    variant="permanent"
                    sx={{
                        display: { xs: 'none', sm: 'block' },
                        '& .MuiDrawer-paper': { boxSizing: 'border-box', width: drawerWidth },
                    }}
                    open
                >
                    {drawer}
                </Drawer>
            </Box>

        </Box>
    );
}

Sidebar.propTypes = {
    /**
     * Injected by the documentation to work in an iframe.
     * You won't need it on your project.
     */
    window: PropTypes.func,
};

export default Sidebar;