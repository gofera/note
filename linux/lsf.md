# submit job
```
$ bsub -q pwodebug "sleep 300"
Job <843569> is submitted to queue <pwodebug>.
```
Submit job in any machines without limit:
```
$ bsub -R "type==any" -q pwodebug "sleep 300" 
```

# list jobs
```
$ bjobs
JOBID   USER    STAT  QUEUE      FROM_HOST   EXEC_HOST   JOB_NAME   SUBMIT_TIME
843569  weliu   RUN   pwodebug   fnode332    fnode438    sleep 300  Apr  1 19:23
841333  weliu   PEND  pwo        dn121201                sleep 500  Apr  1 18:17
```

# kill jobs
```
$ bkill 841333 841334 841338
Job <841333> is being terminated
Job <841334> is being terminated
Job <841338> is being terminated
```
