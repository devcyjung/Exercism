       IDENTIFICATION DIVISION.
       PROGRAM-ID. BINARY-SEARCH.
       ENVIRONMENT DIVISION.
       
       DATA DIVISION.
       WORKING-STORAGE SECTION.
       01 WS-ITEM                  PIC 9999.
       01 WS-RESULT                PIC 99.
       01 WS-ERROR                 PIC X(40).
       01 WS-ARRAY                 PIC X(100).

       01 MID_IDX                  PIC 99.
       01 LEFT_IDX                 PIC 99.
       01 RIGHT_IDX                PIC 99.

       01 ARRAY_LEN                PIC 99.
       01 ARRAY                    PIC 9(4) OCCURS 25.
       
       PROCEDURE DIVISION.
       BINARY-SEARCH. 
           PERFORM PARSE-ARRAY
           MOVE 1 TO LEFT_IDX
           MOVE ARRAY_LEN TO RIGHT_IDX
           MOVE ZERO TO WS-RESULT
           PERFORM UNTIL WS-RESULT > 0 OR LEFT_IDX > RIGHT_IDX
               COMPUTE MID_IDX = (LEFT_IDX + RIGHT_IDX) / 2
               IF WS-ITEM = ARRAY(MID_IDX)
                   MOVE MID_IDX TO WS-RESULT
               END-IF
               IF WS-ITEM < ARRAY(MID_IDX)
                   COMPUTE RIGHT_IDX = MID_IDX - 1
               END-IF
               IF WS-ITEM > ARRAY(MID_IDX)
                   COMPUTE LEFT_IDX = MID_IDX + 1
               END-IF
           END-PERFORM
           IF WS-RESULT = ZERO
               MOVE 'value not in array' TO WS-ERROR
           END-IF.
      
       PARSE-ARRAY.
           MOVE 0 TO ARRAY_LEN
           UNSTRING WS-ARRAY
              DELIMITED BY ","
              INTO ARRAY(1) ARRAY(2) ARRAY(3) ARRAY(4) ARRAY(5)
                   ARRAY(6) ARRAY(7) ARRAY(8) ARRAY(9) ARRAY(10)
                   ARRAY(11) ARRAY(12) ARRAY(13) ARRAY(14) ARRAY(15)
                   ARRAY(16) ARRAY(17) ARRAY(18) ARRAY(19) ARRAY(20)
                   ARRAY(21) ARRAY(22) ARRAY(23) ARRAY(24) ARRAY(25)
              TALLYING IN ARRAY_LEN
           END-UNSTRING.    