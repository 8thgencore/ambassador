import { Button, TextField } from '@mui/material';
import axios from 'axios'
import { SyntheticEvent, useEffect, useState } from 'react'
import Layout from '../components/Layout'


const Profile = (props: any) => {
    const [first_name, setFirstName] = useState("");
    const [last_name, setLastName] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [password_confirm, setPasswordConfirm] = useState("");

    useEffect(() => {
        (
            async () => {
                const { data } = await axios.get("user")

                setFirstName(data.first_name);
                setLastName(data.last_name);
                setEmail(data.email);
            }
        )()
    }, [])

    const infoSubmit = async (e: SyntheticEvent) => {
        // e.preventDefault()

        await axios.put("users/info", {
            first_name,
            last_name,
            email,
        });
    };

    const passwordSubmit = async (e: SyntheticEvent) => {
        // e.preventDefault();

        await axios.put("users/password", {
            password,
            password_confirm,
        });
    };

    return (
        <Layout>
            <h3>Account Information</h3>
            <form onSubmit={infoSubmit}>
                <div className="mb-3">
                    <TextField
                        label="First Name"
                        sx={{ width: '60ch' }}
                        value={first_name}
                        onChange={(e) => setFirstName(e.target.value)}
                    />
                </div>
                <div className="mb-3">
                    <TextField
                        label="Last Name"
                        sx={{ width: '60ch' }}
                        value={last_name}
                        onChange={(e) => setLastName(e.target.value)}
                    />
                </div>
                <div className="mb-3">
                    <TextField
                        value={email}
                        label="Email"
                        sx={{ width: '60ch' }}
                        onChange={(e) => setEmail(e.target.value)}
                    />
                </div>
                <Button variant="contained" className="primary" type="submit">
                    Submit
                </Button>
            </form>

            <h3 className="mt-4">Change Passwords</h3>
            <form onSubmit={passwordSubmit}>
                <div className="mb-3">
                    <TextField
                        label="Password"
                        type="password"
                        sx={{ width: '60ch' }}
                        onChange={(e) => setPassword(e.target.value)}
                    />
                </div>
                <div className="mb-3">
                    <TextField
                        label="Password Confirm"
                        type="password"
                        sx={{ width: '60ch' }}
                        onChange={(e) => setPasswordConfirm(e.target.value)}
                    />
                </div>
                <Button variant="contained" className="primary" type="submit">
                    Submit
                </Button>
            </form>
        </Layout>
    );
};

export default Profile;