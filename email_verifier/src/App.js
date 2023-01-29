import './App.css';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';

function App() {
  return (
    <div className="App">
    
      <div class="head">
      <Typography variant="h2">
        XERO MAIL
      </Typography>
      </div>
    
      <div class="subs">
    <Typography variant="subtitle1">
    A free e-mail checking tool to check your address domains and verify its presence.
    </Typography>  
    </div>

    <div class="entry">
      <div class="entryfield">
    <TextField id="outlined-basic" label="Outlined" variant="outlined" />
    </div>
    <div class="entrybutton">
    <Button variant="outlined">Outlined</Button>
    </div>
    </div>
    
    <div class="footer">
      <div class="foot1">
      <Typography variant="h4">
      Good outreach is only possible if you reach the lead’s inbox.
    </Typography>
    </div>  
    <div class="foot2">
    <Typography variant="subtitle1">
    Check every email address you have and reduce your bounce rate.
    </Typography>
    </div>

    </div>

    </div>
  );
}

export default App;
