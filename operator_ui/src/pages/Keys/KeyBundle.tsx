import React from 'react'
import Typography from '@material-ui/core/Typography'
import VpnKeyIcon from '@material-ui/icons/VpnKey'
import List from '@material-ui/core/List'
import ListItem from '@material-ui/core/ListItem'
import ListItemIcon from '@material-ui/core/ListItemIcon'
import ListItemText from '@material-ui/core/ListItemText'
import Avatar from '@material-ui/core/Avatar'
import {
  createStyles,
  withStyles,
  WithStyles,
  Theme,
} from '@material-ui/core/styles'

const styles = (theme: Theme) =>
  createStyles({
    listItemPrimary: {
      marginBottom: theme.spacing.unit,
    },
  })

export const KeyBundle = withStyles(styles)(
  ({
    classes,
    primary,
    secondary,
  }: WithStyles<typeof styles> & {
    primary: React.ReactNode
    secondary: React.ReactNode[]
  }) => {
    return (
      <List dense={true}>
        <ListItem>
          <ListItemIcon>
            <Avatar>
              <VpnKeyIcon />
            </Avatar>
          </ListItemIcon>
          <ListItemText
            primary={
              <Typography
                className={classes.listItemPrimary}
                variant="body1"
                component="span"
              >
                {primary}
              </Typography>
            }
            secondary={secondary.map((secondaryItem: any, index: number) => (
              <Typography key={index} variant="h6" component="span">
                {secondaryItem}
              </Typography>
            ))}
          />
        </ListItem>
      </List>
    )
  },
)
