import { createUserWithEmailAndPassword } from "firebase/auth"
import { useState } from "react"
import { useNavigate } from "react-router-dom";
import { firebaseAuth } from "../firebase_utils/auth"

export const CreateAccount: React.FC = () => {
  const [isPending, setIsPending] = useState(false);
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [errMsg, setErrorMsg] = useState("");

  const navigate = useNavigate();

  const createAccountWithEmailAndPassword = () => {
    setErrorMsg("");
    setIsPending(true);
    createUserWithEmailAndPassword(firebaseAuth, email, password)
      .then(()=>{
        navigate('/');
      })
      .catch((reason) => {
        setIsPending(false);
        setErrorMsg(reason.toString());
      });
  };


  return (
    <div>
      <input placeholder="メールアドレス" type="text" onChange={(e) => {setEmail(e.target.value)}}></input>
      <input placeholder="パスワード" type="text" onChange={(e) => {setPassword(e.target.value)}}></input>
      <button onClick={() => {
        if (!isPending)createAccountWithEmailAndPassword();
      }}
      >アカウントを作成</button>
      {isPending ? <div>アカウントを作成中...</div> : null}
      <div>{errMsg}</div>
    </div>
  ) 
}