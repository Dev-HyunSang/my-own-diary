# my own diary
많은 분들께서 다이어리를 쓰시는 모습을 보고 "다이어리만 쓸 수 있는 플랫폼이 있으면 어떨까?"라는 생각이 들어서 개발하게 되었습니다.    
"**나의 생각**"을 담는 서비스인 만큼 보안을 신경 써서 개발하고 있습니다.  

## TODO:
- [X] POST `/register`: 회원가입 - 2022.03.12 개발 완료
- [ ] POST `/login`: 로그인
- [ ] POST `/logout`: 로그아웃

### User
- [ ] POST`/user/edit`: 회원정보 수정

## 사용한 기술 스택
> 사용한 기술 스택에 대해서 서술하였습니다. 

- **Back-End: Golang**
  - [`gofiber/fiber`](https://gofiber.io/)
  - [`gorm/gorm`](gorm.io/gorm)
- DataBase: MySQL