import './App.css';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';

function App() {
  return (
    <div className="App">
      <Typography variant="h2">
        XERO MAIL
    </Typography>
    <Typography variant="subtitle1">
    A free e-mail checking tool to check your address domains and verify its presence.
    </Typography>  
    <TextField id="outlined-basic" label="Outlined" variant="outlined" />
    <Button variant="outlined">Outlined</Button>
    <Typography variant="h4">
    Good outreach is only possible if you reach the lead’s inbox.
    </Typography>  
    <Typography variant="subtitle1">
    Check every email address you have and reduce your bounce rate.
    </Typography>  
    </div>
  );
}

export default App;
