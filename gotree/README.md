## gotree tool ##

### how to install 
1. download the source code 
2. go build 
3. go install 

### how to use

gotree 
显示所有内容
~~~ bash
laibr-a@laibr-a-compute:~/golang/workspace/src/minitask$ gotree 
 |-- golangsmalltool 
    |-- .git 
       |-- logs 
          |-- refs 
             |-- heads 
                |-- master
             |-- remotes 
                |-- origin 
                   |-- master
                   |-- HEAD
          |-- HEAD
       |-- packed-refs
       |-- branches 
       |-- refs 
          |-- tags 
          |-- heads 
             |-- master
          |-- remotes 
             |-- origin 
                |-- master
                |-- HEAD
       |-- objects 
          |-- cd 
             |-- b2839fb6ba95c5365381bfb5a33d4947ea9ba9
          |-- 6b 
             |-- c1a1f131a9c0c017af592bf501618c034efc3c

~~~

gotree --maxlevel 1
显示一层目录
~~~ bash
laibr-a@laibr-a-compute:~/golang/workspace/src/minitask$ gotree --maxlevel 1
 |-- golangsmalltool 
    |-- .git 
    |-- MyStruct.go
    |-- goStruct Note.md
    |-- CommandServer.go
    |-- CommandClient.go
    |-- gotree 
    |-- README.md
    |-- Hello.go


~~~