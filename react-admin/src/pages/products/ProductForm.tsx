import { Button, TextField } from '@mui/material'
import axios from 'axios';
import { SyntheticEvent, useEffect, useState } from 'react'
import { Navigate, useParams } from 'react-router-dom';
import Layout from '../../components/Layout'

const ProductForm = () => {
    const [title, setTitle] = useState("");
    const [description, setDescription] = useState("");
    const [image, setImage] = useState("");
    const [price, setPrice] = useState("");
    const [redirect, setRedirect] = useState(false);


    const { id } = useParams();

    useEffect(() => {
        if (id) {
            (async () => {
                const { data } = await axios.get(`products/${id}`);

                setTitle(data.title);
                setDescription(data.description);
                setImage(data.image);
                setPrice(data.price);
            })();
        }
    }, []);

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();
        const data = {
            title,
            description,
            image,
            price: parseFloat(price),
        };

        if (id) {
            await axios.put(`products/${id}`, data);
        } else {
            await axios.post("products", data);
        }

        setRedirect(true);
    };

    if (redirect) {
        return <Navigate to={"/products"} />;
    }

    return (
        <Layout>
            <form onSubmit={submit}>
                <h2>Add new product</h2>
                <div className="mb-3">
                    <TextField
                        label="Title"
                        sx={{ width: '60ch' }}
                        value={title}
                        onChange={(e: any) => setTitle(e.target.value)}
                    />
                </div>
                <div className="mb-3">
                    <TextField
                        label="Description"
                        rows={4}
                        multiline
                        sx={{ width: '60ch' }}
                        value={description}
                        onChange={(e) => setDescription(e.target.value)}
                    />
                </div>
                <div className="mb-3">
                    <TextField
                        value={image}
                        label="Image"
                        sx={{ width: '60ch' }}
                        onChange={(e) => setImage(e.target.value)}
                    />
                </div>
                <div className="mb-3">
                    <TextField
                        value={price}
                        label="Price"
                        type="number"
                        sx={{ width: '60ch' }}
                        onChange={(e) => setPrice(e.target.value)}
                    />
                </div>
                <Button variant="contained" color="primary" type="submit">
                    Submit
                </Button>
            </form>
        </Layout>
    )
}

export default ProductForm